import { Component, ElementRef, Input, OnInit, ViewChild } from '@angular/core';
import { Vote } from 'src/app/interfaces/vote';
import { FeaturesService } from 'src/app/repositories/features.service';
import { NzModalService } from 'ng-zorro-antd/modal';
import { UserModalData } from 'src/app/interfaces/user';
import { CommentDetailComponent } from '../comment-detail/comment-detail.component';
import { Comment } from 'src/app/interfaces/comment';
import { environment } from 'src/environments/environment';
import { ActivatedRoute, Router } from '@angular/router';
import { NzImageService } from 'ng-zorro-antd/image';
import { ApiService } from 'src/app/services/api.service';
import { LoaderService } from 'src/app/services/loader.service';
import { BehaviorSubject } from 'rxjs';
import { MasterService } from 'src/app/services/master.service';
import { User } from '../../interfaces/user';

@Component({
  selector: 'app-list-table',
  templateUrl: './list-table.component.html',
  styleUrls: ['./list-table.component.scss']
})
export class ListTableComponent implements OnInit {
  @ViewChild('ellipsisContainer') ellipsisContainer!: ElementRef;
  private _data = new BehaviorSubject<Vote[]>([]);
  dataList!: Vote[];

  currentPage = 1;
  // @Input() data: Vote[] = [];
  @Input() type!: string
  @Input() summaryCount!: number;
  @Input()
  set data(value) { this._data.next(value); }
  get data() { return this._data.getValue(); }
  loading = true;
  baseUrl = environment.angularBaseUrl;
  userId!: string;
  isLoading = false;
  userData: User[] = [];

  constructor(
    private featureSvc: FeaturesService,
    private modal: NzModalService,
    private nzImageService: NzImageService,
    private router: Router,
    private apiSvc: ApiService,
    private loaderSvc: LoaderService,
    private route: ActivatedRoute,
    private masterSvc: MasterService
  ) { }

  ngOnInit(): void {
    this.masterSvc.userData$.subscribe((userData: User[]) => {
      this.userData = userData;
    });
    this._data.subscribe((data: Vote[]) => {
      this.dataList = data;
      if (data) {
        this.dataList = this.dataList.map((data: Vote) => {
          const user = this.userData.find(user => user._id === data.user_id);
          return {
            ...data,
            user_profile: user?.profile,
            display_name: user?.display_name
          }
        });
      }
      this.loading = this.dataList ? false : true;
    });
    this.route.paramMap.subscribe(params => {
      this.userId = params.get('id') || '';
    });
  }

  loadMore(): void {
    this.currentPage++;
    const limit = 10;
    this.isLoading = true;
    if (this.type === 'vote') {
      if (this.userId != null) {
        this.featureSvc.getVotesByUserId(this.userId, this.currentPage, limit).subscribe({
          next: (data: Vote[]) => {
            this.data = [...this.data, ...data];
            this.isLoading = false;
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.apiSvc.handleMessageError(error);
            this.isLoading = false;
          }
        })
      }
    }
    if (this.type === 'bookmark') {
      if (this.userId != null) {
        this.featureSvc.getBookmarkByUserId(this.userId, this.currentPage, limit).subscribe({
          next: (data: Vote[]) => {
            this.data = [...this.data, ...data];
            this.isLoading = false;
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.apiSvc.handleMessageError(error);
            this.isLoading = false;
          }
        })
      }
    }
    if (this.type === 'answers') {
      if (this.userId != null) {
        this.featureSvc.getAnswersByUserId(this.userId, this.currentPage, limit).subscribe({
          next: (data: Vote[]) => {
            this.data = [...this.data, ...data];
            this.isLoading = false;
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.apiSvc.handleMessageError(error);
            this.isLoading = false;
          }
        })
      }
    }
    if (this.type === 'solved') {
      if (this.userId != null) {
        this.featureSvc.getCommentByUserIdSolved(this.userId, this.currentPage, limit).subscribe({
          next: (data: Vote[]) => {
            this.data = [...this.data, ...data];
            this.isLoading = false;
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.apiSvc.handleMessageError(error);
            this.isLoading = false;
          }
        })
      }
    }
    if (this.type === 'mention') {
      if (this.userId != null) {
        this.featureSvc.getMentionByUserId(this.userId, this.currentPage, limit).subscribe({
          next: (data: Vote[]) => {
            this.data = [...this.data, ...data];
            this.isLoading = false;
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.apiSvc.handleMessageError(error);
            this.isLoading = false;
          }
        })
      }
    }
  }

  removeHtmlTags(input: string) {
    if (input) {
      return input.replace(/<[^>]*>/g, '').replace(/&[^;]+;/g, '');
    }
    return
  }

  Show(daa: any): void {
    const data: any = {
      type: 'commentDetail',
      description: daa.description
    }
    this.showDetail(data);
  }

  showDetail(nzData: UserModalData) {
    this.modal.create({
      nzContent: CommentDetailComponent,
      nzClosable: false,
      nzCentered: true,
      nzMaskClosable: false,
      nzStyle: { 'width': '80vw' },
      nzData: nzData,
    })
  }

  viewComment(comment: Comment) {
    const link = this.getCommentLink(comment);
    window.open(link, '_self');
  }

  getCommentLink(comment: Comment): string {
    return `/qa-detail/${comment.question_id}#${comment._id}`;
  }

  previewImg(src: string | undefined): void {
    const images = [{
      src: src ? src : '../../../assets/images/avatar.png'
    }];
    this.nzImageService.preview(images, { nzZoom: 1, nzRotate: 0 });
  }
}
