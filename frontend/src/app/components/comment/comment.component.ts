import { Component, OnInit, inject } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { NzModalRef, NZ_MODAL_DATA, NzModalService } from 'ng-zorro-antd/modal';
import { User, UserModalData } from '../../interfaces/user';
import { RequestComment } from '../../interfaces/comment';
import { CommentService } from '../../repositories/comment.service';
import { FeaturesService } from 'src/app/repositories/features.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { UserService } from 'src/app/repositories/user.service';
import { TranslateService } from '@ngx-translate/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { LoaderService } from 'src/app/services/loader.service';
import { AuthService } from 'src/app/services/auth.service';
import { Comment } from '../../interfaces/comment';
import { ApiService } from 'src/app/services/api.service';
import { MasterService } from 'src/app/services/master.service';

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.scss']
})
export class CommentComponent implements OnInit {
  isLoading = false;
  commentForm!: UntypedFormGroup;
  msg!: string;
  imageUrl!: string;
  cloudinaryUrlDesc!: string;
  showToolbar = true;
  disabled = false;
  public readonly = true;
  readonly userModalData: UserModalData = inject(NZ_MODAL_DATA);
  userId!: string | null;
  isEdited = false;
  userData: User[] = [];

  constructor(
    private fb: UntypedFormBuilder,
    public modal: NzModalService,
    private modalRef: NzModalRef,
    private message: NzMessageService,
    private commentSvc: CommentService,
    private featureSvc: FeaturesService,
    private userSvc: UserService,
    private translate: TranslateService,
    private messageSvc: NzMessageService,
    private http: HttpClient,
    private loaderSvc: LoaderService,
    private authSvc: AuthService,
    private apiSvc: ApiService,
    private masterSvc: MasterService
  ) { }

  ngOnInit(): void {
    this.userId = localStorage.getItem('login_user_id') || ''
    this.commentForm = this.fb.group({
      description: [null, [Validators.required]],
    });

    if (this.userModalData.type === 'update') {
      this.commentForm.patchValue({
        description: this.userModalData.description
      })
    }
    this.masterSvc.userData$.subscribe((userData: User[]) => {
      this.userData = userData;
    });
  }

  cancel(): void {
    this.modalRef.destroy("cancel");
  }

  submit(): void {
    const description = this.commentForm.get('description')?.value;
    const srcRegex = /<img[^>]+src="([^">]+)"/i;
    if (this.userModalData.type == "create") {
      this.msg = this.translate.instant('commentCreate')
    } else {
      this.msg = this.translate.instant('commentUpdate')
    }
    const type = this.userModalData.type == 'create' ? this.translate.instant('createComment') : this.translate.instant('updateComment')
    this.modal.confirm({
      nzTitle: `<i>${type}</i>`,
      nzContent: `<b>${this.msg}</b>`,
      nzOnOk: () => {
        if (this.userModalData.type === 'create') {
          if (srcRegex.test(description)) {
            this.changeImageUrl('create')
          } else {
            this.createComment()
          }
        } else if (this.userModalData.type == 'update') {
          if (srcRegex.test(description)) {
            this.changeImageUrl('update')
          } else {
            this.updateComment()
          }
        }
      },
      nzOkText: this.translate.instant("ok"),
      nzCancelText: this.translate.instant("cancel")
    })
  }

  getMentionsFromBody(description: string): string[] {
    const mentionRegex = /data-mention="@(.*?)"/g;
    const mentions = description.match(mentionRegex);
    const uniqueMentions = new Set(mentions?.map((mention: any) => mention.replace('data-mention="@', '').replace('"', '')) || []);
    return Array.from(uniqueMentions);
  }

  handleMentions(data: any, mentionedUsernames: string[]): void {
    
    if (mentionedUsernames.length > 0) {
      mentionedUsernames.map((displayName) => {
        this.userSvc.getUserByDisplayName(displayName).subscribe({
          next: (user) => {
            let tokenId = this.getUserNotiToken(user._id)
            const mentionData = {
              user_id: user._id,
              question_id: data.question_id,
              comment_id: data._id,
              comment_link: "link",
              noti_token: tokenId
            };
            this.featureSvc.createMention(mentionData).subscribe({
              next: () => {
                //.....//
              }
            });
          }
        });
      });
    } else {
      console.log("Description does not contain a mention.");
    }
  }

  uploadImage(file: File) {
    const uploadData = new FormData();
    uploadData.append('file', file);
    uploadData.append('upload_preset', environment.preset);
    uploadData.append('folder', 'mtm_community_profile');
    return this.http.post(`https://api.cloudinary.com/v1_1/${environment.cloudName}/upload?upload_preset=${environment.preset}`, uploadData);
  }

  changeImageUrl(cond: string) {
    const description = this.commentForm.get('description')?.value;
    const srcRegex = /<img[^>]+src="([^">]+)"/i;
    const match = description.match(srcRegex);
    if (match) {
      const imageSrc = match[1];
      const isBase64 = this.isBase64Image(imageSrc)
      if (isBase64) {
        this.uploadImage(imageSrc).subscribe({
          next: (response: any) => {
            this.imageUrl = response.url
            const img = description.replace(srcRegex, `<img src="${this.imageUrl}"`);
            this.commentForm.controls['description'].setValue(img);
            if (cond == 'create') {
              this.createComment();
            } else if (cond == 'update') {
              this.updateComment();
            }

            this.messageSvc.success(this.translate.instant('imageUpload'));
          },
          error: () => {
            this.messageSvc.error(this.translate.instant('wrong'));
          }
        })
      } else {
        if (cond == 'create') {
          this.createComment();
        } else if (cond == 'update') {
          this.updateComment();
        }
      }
    }
  }

  isBase64Image(imageSrc: string): boolean {
    const base64PrefixRegex = /^data:image\/[^;]+;base64,/i;
    const base64Data = imageSrc.replace(base64PrefixRegex, '');
    const isValidBase64 = /^[a-z0-9+/]+={0,2}$/i.test(base64Data);
    return isValidBase64;
  }

  createComment(): void {
    this.isLoading = true;
    let tokenId = this.getUserNotiToken(this.userModalData.question_detail_user_id || '')
    const body: RequestComment = {
      user_id: this.userModalData.user_id,
      parent_id: this.userModalData.parent_id,
      question_id: this.userModalData.question_id,
      sort: this.userModalData.sort,
      description: this.commentForm.get('description')?.value,
      noti_token: tokenId
    };
    if (this.commentForm.valid) {
      this.loaderSvc.call();
      this.commentSvc.createComment(body).subscribe({
        next: (data) => {
          const mentionedUsernames = this.getMentionsFromBody(body.description);
          this.commentSvc.getComment(data.parent_id).subscribe({
            next: () => {
              this.loaderSvc.dismiss();
            },
            error: (error) => {
              this.loaderSvc.dismiss();
              console.log(error)
            }
          })
          this.handleMentions(data, mentionedUsernames)
          this.isLoading = false;
          this.message.create('success', this.translate.instant('commentCreateSuccess'));
          this.modalRef.destroy('success');
        },
        error: (error) => {
          this.isLoading = false;
          this.loaderSvc.dismiss();
          this.apiSvc.handleMessageError(error);
          this.apiSvc.handleErrorType(error);
        }
      });
    } else {
      this.messageSvc.error(this.translate.instant('invalidFormData'));
      this.loaderSvc.dismiss();
      this.isLoading = false;
    }
  }

  updateComment(): void {
    this.isLoading = true;
    this.isEdited = true;
    const body: RequestComment = {
      user_id: this.userModalData.user_id,
      parent_id: this.userModalData.parent_id,
      question_id: this.userModalData.question_id,
      sort: this.userModalData.sort,
      description: this.commentForm.get('description')?.value,
    };

    if (this.commentForm.valid) {
      this.loaderSvc.call();
      this.commentSvc.updateComment(this.userModalData.comment_id || '', body).subscribe({
        next: (data) => {
          const mentionedUsernames = this.getMentionsFromBody(body.description);
          this.commentSvc.getComment(data._id).subscribe({
            next: (commentData: Comment) => {
              if (commentData.user_id === this.userId) {
                this.authSvc.noti.next(this.userModalData.count + 1);
              }
              this.loaderSvc.dismiss();
            },
            error: () => {
              this.loaderSvc.dismiss();
            }
          })
          this.handleMentions(data, mentionedUsernames);
          this.isLoading = false;
          this.message.create('success', this.translate.instant('commentUpdateSuccess'));
          this.modalRef.destroy('success');
        },
        error: (error) => {
          this.isLoading = false;
          this.loaderSvc.dismiss();
          this.apiSvc.handleMessageError(error);
          this.apiSvc.handleErrorType(error);
        }
      });
    } else {
      this.messageSvc.error(this.translate.instant('invalidFormData'));
      this.loaderSvc.dismiss();
      this.isLoading = false;
    }
  }

  getUserNotiToken(id: string): string {
    return this.userData.find((user: User) => user._id === id)?.noti_token || '';
  }
}

