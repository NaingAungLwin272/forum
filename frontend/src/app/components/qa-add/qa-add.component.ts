import { Component, Input, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { NzModalService } from 'ng-zorro-antd/modal';
import { NzDrawerRef } from 'ng-zorro-antd/drawer';
import { Category } from '../../interfaces/category';
import { User } from '../../interfaces/user';
import { QuestionService } from 'src/app/repositories/question.service';
import { CommentService } from 'src/app/repositories/comment.service';
import { NzMessageService } from "ng-zorro-antd/message";
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { TranslateService } from '@ngx-translate/core';
import { UserService } from 'src/app/repositories/user.service';
import { FeaturesService } from 'src/app/repositories/features.service';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';
import { MasterService } from 'src/app/services/master.service';

@Component({
  selector: 'app-qa-add',
  templateUrl: './qa-add.component.html',
  styleUrls: ['./qa-add.component.scss']
})
export class QaAddComponent implements OnInit {
  @Input() categoryData!: Category[];
  @Input() userData!: User[];
  @Input() questionCount!: number;
  qaAddForm!: UntypedFormGroup;
  isLoading = false;
  imageUrl!: string;
  cloudinaryUrlDesc!: string;
  readonly = true;
  lastpost!: string;
  userId!: string;
  isCheckedButton = false;
  User = {} as User


  constructor(
    private fb: UntypedFormBuilder,
    public modal: NzModalService,
    private drawerRef: NzDrawerRef<string>,
    private QuestionSvc: QuestionService,
    private messageSvc: NzMessageService,
    private commentSvc: CommentService,
    private userSvc: UserService,
    private featureSvc: FeaturesService,
    private http: HttpClient,
    private translate: TranslateService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService,
    private masterSvc: MasterService
  ) {
    this.userId = localStorage.getItem('login_user_id') || ''
  }

  ngOnInit(): void {
    this.qaAddForm = this.fb.group({
      user_id: [this.userId],
      title: [null, [Validators.required]],
      language_ids: [[], [Validators.required]],
      tag_ids: [[], [Validators.required]],
      description: [null, [Validators.required]],
      user_ids: [[this.userId]],
      is_mentioned: [],
      user_name: [],
    });

    // bind if data exist
    const question = localStorage.getItem('question');
    if (question) {
      this.qaAddForm.patchValue(JSON.parse(question));
    }


    // save data when x close
    this.drawerRef.afterClose.subscribe(data => {
      if (!data) {
        localStorage.setItem('question', JSON.stringify(this.qaAddForm.value));
      }
    })

    this.userSvc.getUser(this.userId).subscribe({
      next:(loginUserData: User) => {
        this.qaAddForm.value.user_name = loginUserData.name
      }
    })
  }

  uploadImage(file: File) {
    const uploadData = new FormData();
    uploadData.append('file', file);
    uploadData.append('upload_preset', environment.preset);
    uploadData.append('folder', 'mtm_community_profile');
    return this.http.post(`https://api.cloudinary.com/v1_1/${environment.cloudName}/upload?upload_preset=${environment.preset}`, uploadData);
  }

  changeImageUrl() {
    const description = this.qaAddForm.get('description')?.value;
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
            this.qaAddForm.controls['description'].setValue(img);
            this.questionCreate();
            this.messageSvc.success(this.translate.instant('imageUpload'));
          },
          error: () => {
            this.messageSvc.error(this.translate.instant('wrong'));
          }
        })
      } else {
        this.messageSvc.error(this.translate.instant('notBase64'));
      }
    }
  }

  isBase64Image(imageSrc: string): boolean {
    const base64PrefixRegex = /^data:image\/[^;]+;base64,/i;
    const base64Data = imageSrc.replace(base64PrefixRegex, '');
    const isValidBase64 = /^[a-z0-9+/]+={0,2}$/i.test(base64Data);
    return isValidBase64;
  }

  getMentionsFromBody(description: string): string[] {
    const mentionRegex = /data-mention="@(.*?)"/g;
    const mentions = description.match(mentionRegex);
    return mentions?.map((mention: any) => mention.replace('data-mention="@', '').replace('"', '')) || [];
  }

  handleMentions(data: any, commentId: string, mentionedUsernames: string[]): void {
    if (mentionedUsernames.length > 0) {
      mentionedUsernames.map((displayName) => {
        this.loaderSvc.call();
        this.userSvc.getUserByDisplayName(displayName).subscribe({
          next: (user) => {
            const mentionData = {
              user_id: user._id,
              question_id: data.QuestionRes._id,
              comment_id: commentId,
            };
            this.featureSvc.createMention(mentionData).subscribe({
              next: () => {
                this.loaderSvc.dissAll();
              }
            });
          }
        });
      });
    } else {
      console.log("Description does not contain a mention.");
      this.loaderSvc.dismiss();
    }
  }

  async submitForm(): Promise<void> {
    const questionConfirmText = this.translate.instant('questionCreateConfirm');
    const questionCreate = this.translate.instant('questionCreateInfo')
    const description = this.qaAddForm.get('description')?.value;
    const srcRegex = /<img[^>]+src="([^">]+)"/i;
    this.modal.confirm({
      nzTitle: `<i>${questionCreate}</i>`,
      nzContent: `<b>${questionConfirmText}</b>`,
      nzOnOk: async () => {
        this.isLoading = true;
        if (srcRegex.test(description)) {
          this.changeImageUrl();
        } else {
          this.questionCreate();
        }
      },
      nzOkText: this.translate.instant("ok"),
      nzCancelText: this.translate.instant("cancel")
    })
  }

  questionCreate() {
    if (this.qaAddForm.valid) {
      this.qaAddForm.value.is_mentioned = this.isCheckedButton
      this.loaderSvc.call();
      this.QuestionSvc.createQuestion(this.qaAddForm.value).subscribe({
        next: (data) => {
          const mentionedUsernames = this.getMentionsFromBody(this.qaAddForm.get('description')?.value);
          const commentId = data.CommentId;
          this.handleMentions(data, commentId, mentionedUsernames);
          this.isLoading = false;
          this.messageSvc.success(this.translate.instant('questionCreate'));
          localStorage.removeItem('question');
          this.drawerRef.close("created");
          this.loaderSvc.dismiss();
        },
        error: (error) => {
          this.loaderSvc.dismiss();
          this.apiSvc.handleMessageError(error);
          this.apiSvc.handleErrorType(error);
          this.drawerRef.close();
          this.isLoading = false;
        }
      });
    } else {
      this.messageSvc.error(this.translate.instant('invalidFormData'));
      this.loaderSvc.dismiss();
      this.isLoading = false;
    }
  }
}
