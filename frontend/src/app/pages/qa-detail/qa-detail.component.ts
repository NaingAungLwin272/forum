import * as customBuild from '../../components/ckCus/build/ckeditor';

import { AfterViewInit, ChangeDetectorRef, Component, ElementRef, OnInit } from '@angular/core';
import { Bookmark, CreateBookmark } from '../../interfaces/bookmark';
import { Comment, RequestComment } from 'src/app/interfaces/comment';
import { CreateUserVote, UserVote } from '../../interfaces/vote';
import { User, UserModalData } from '../../interfaces/user';

import { ActivatedRoute } from '@angular/router';
import { ApiService } from 'src/app/services/api.service';
import { AuthService } from 'src/app/services/auth.service';
import { BadgesService } from 'src/app/repositories/badges.service';
import { Category } from 'src/app/interfaces/category';
import { CategoryService } from 'src/app/repositories/category.service';
import { CommentComponent } from '../../components/comment/comment.component';
import { CommentService } from '../../repositories/comment.service';
import { FeaturesService } from 'src/app/repositories/features.service';
import { LoaderService } from 'src/app/services/loader.service';
import { MasterService } from 'src/app/services/master.service';
import { NzImageService } from 'ng-zorro-antd/image';
import { NzMessageService } from "ng-zorro-antd/message";
import { NzModalService } from 'ng-zorro-antd/modal';
import { QuestionDetail } from 'src/app/interfaces/question';
import { QuestionService } from '../../repositories/question.service';
import { Router } from '@angular/router';
import { ThemeService } from 'src/app/services/theme.service';
import { TranslateService } from '@ngx-translate/core';
import { UserService } from '../../repositories/user.service';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-qa-detail',
  templateUrl: './qa-detail.component.html',
  styleUrls: ['./qa-detail.component.scss']
})
export class QaDetailComponent implements OnInit, AfterViewInit {
  isOpened !: boolean;
  userData: User[] = [];
  questionDetail!: QuestionDetail;
  commentList!: Comment[];
  categoryData: Category[] = [];
  qaId!: string;
  darkTheme!: boolean;
  baseUrl = environment.angularBaseUrl;
  private scrollToCommentId: string | null = null;
  userId = localStorage.getItem('login_user_id');
  notiCount = 0;
  isClicked !: boolean;
  solutionClicked !: boolean;
  bookmarkClicked !: boolean;
  uniqueUserIds = new Set<string>();
  commentUserIds: Map<string, Set<string>> = new Map();
  solutionCount: any;
  voteCount: any;
  replyCount!: number;
  solutionComment: any;
  showAllAvatars = true;
  activePanels: string[] = [];
  firstHashed!: string;
  example !: boolean;
  editor = customBuild;
  targetEle!: any;
  editorConfig = {
    toolbar: {
      items: [],
      shouldNotGroupWhenFull: true,
    },
  };

  userVoteData: UserVote[] = [];
  bookmarkData: Bookmark[] = [];

  constructor(
    private route: ActivatedRoute,
    public modal: NzModalService,
    private userSvc: UserService,
    private questionSvc: QuestionService,
    private categorySvc: CategoryService,
    private featureSvc: FeaturesService,
    private themeService: ThemeService,
    private messageSvc: NzMessageService,
    private elementRef: ElementRef,
    private nzImageService: NzImageService,
    private commentSvc: CommentService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService,
    private authSvc: AuthService,
    private badgeSvc: BadgesService,
    private translate: TranslateService,
    private router: Router,
    private masterSvc: MasterService,
    private cdRef: ChangeDetectorRef
  ) { }

  ngOnInit(): void {
    this.qaId = this.route.snapshot.paramMap.get('id') || '';
    this.darkTheme = localStorage.getItem('theme') === 'dark' ? true : false;
    enum ThemeType {
      dark = 'dark',
      default = 'default',
    }
    this.themeService.currentTheme$.subscribe((currentTheme: ThemeType) => {
      this.darkTheme = currentTheme === ThemeType.dark;
    });
    this.getCommentFragment();
    this.getQaDetail();
    this.authSvc.noti$.subscribe((data: number) => {
      this.notiCount = data;
    });
    this.masterSvc.userData$.subscribe((userData: User[]) => {
      this.userData = userData;
    });
    this.masterSvc.categoryData$.subscribe((categoryData: Category[]) => {
      this.categoryData = categoryData;
    });
  }

  ngAfterViewInit(): void {
    if (this.questionDetail) {
      this.scrollCommon();
    }

  }

  scrollCommon(): void {
    if (this.scrollToCommentId) {
      const parentId = this.findParentId(this.commentList, this.scrollToCommentId);
      if (parentId) {
        this.openCollapsePanel(parentId);
      }

      const target = this.elementRef.nativeElement.querySelector(`[id="${this.scrollToCommentId}"]`);
      if (target) {
        setTimeout(() => {
          target.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }, 500);
      }
    }
  }

  findParentId(comments: Comment[], targetCommentId: string): string | null {
    for (const comment of comments) {
      if (comment.reply_comments && comment.reply_comments.length > 0) {
        const parentId = this.findParentId(comment.reply_comments, targetCommentId);
        if (parentId) {
          return parentId;
        }
      }
      if (comment._id === targetCommentId) {
        return comment.parent_id;
      }
    }
    return null;
  }

  getCommentFragment() {
    const hashFragment = window.location.hash;
    if (hashFragment) {
      const commentId = hashFragment.substring(1);
      this.scrollToCommentId = commentId;
    }
  }

  handleAnchorClick(commentId: string): void {
    const indexAfterHashQA = commentId.indexOf("#qa") + "#qa".length;
    const splitedCommentId = commentId.substring(indexAfterHashQA);
    if (!this.activePanels.includes(splitedCommentId)) {
      this.activePanels.push(splitedCommentId);
    }
  }

  getQaDetail() {
    this.questionSvc.getQuestionById(this.qaId).subscribe({
      next: (questionDetail: QuestionDetail) => {
        this.questionDetail = questionDetail;
        this.getBookmarkByUserIdQuestionId();
      },
      error: (error) => {
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
        this.router.navigate(['/']);
      }
    });
  }

  getBookmarkByUserIdQuestionId() {
    if (this.userId != null) {
      this.featureSvc.getBookmarkByUserIdQuestionId(this.userId, this.questionDetail._id).subscribe({
        next: (bookmarkData: Bookmark[]) => {
          this.bookmarkData = bookmarkData || [];
          this.getVotesByUserIdQuestionId();
        }
      });
    }
  }

  getVotesByUserIdQuestionId() {
    this.loaderSvc.call();
    if (this.userId != null) {
      this.featureSvc.getVotesByUserIdQuestionId(this.userId, this.questionDetail._id).subscribe({
        next: (userVoteData: UserVote[]) => {
          this.userVoteData = userVoteData || [];
          this.commentList = this.generateCommentList(this.questionDetail.comments);
          this.populateUniqueUserIds();
          this.commentList.forEach((comment, index) => {
            if (index === this.commentList.length - 1) {
              this.loaderSvc.dissAll();
              this.scrollAfterCommentsLoaded();
            }
          });
          this.loaderSvc.dissAll();
        },
        error: () => {
          this.loaderSvc.dissAll();
        }
      });
    }
  }

  scrollAfterCommentsLoaded() {
    setTimeout(() => {
      this.scrollCommon();
    }, 0);
  }

  openCollapsePanel(commentId: string) {
    if (!this.activePanels.includes(commentId)) {
      this.activePanels.push(commentId);
      const target = this.elementRef.nativeElement.querySelector(`[id="${this.scrollToCommentId}"]`);
      if (target) {
        setTimeout(() => {
          target.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }, 1000);
      }
    }
  }

  activeChange(active: boolean, commentId: string) {
    this.example = active;
    if (!active) {
      this.closePanel(commentId);
    } else {
      this.openPanel(commentId);
    }
  }

  openPanel(commentId: string) {
    if (!this.activePanels.includes(commentId)) {
      this.activePanels.push(commentId);
      this.cdRef.detectChanges();
    }
  }

  closePanel(commentId: string) {
    const index = this.activePanels.indexOf(commentId);
    if (index !== -1) {
      this.activePanels.splice(index, 1);
      this.cdRef.detectChanges();
    }
  }

  openCollapseAndNavigate(commentId: string, link: string) {
    this.openCollapsePanel(commentId);
    const target = this.elementRef.nativeElement.querySelector(`[id="${commentId}"]`);
    if (target) {
      target.scrollIntoView({ behavior: 'smooth', block: 'start' });
      setTimeout(() => {
        window.location.href = link;
      }, 500);
    } else {
      window.location.href = link;
    }
  }

  isPanelActive(commentId: string): boolean {
    return this.activePanels.includes(commentId);
  }

  generateCommentList(array: Comment[]) {
    if (!Array.isArray(array) || array.length === 0) {
      return [];
    }

    const tree: Comment[] = [];

    for (let i = 0; i < array.length; i++) {
      array[i].solutionCount = 0;
      array[i].voteCount = 0;

      if (array[i].parent_id) {
        const parent: Comment | undefined = array.find(elem => elem._id === array[i].parent_id);
        if (parent) {
          if (!parent.reply_comments) {
            parent.reply_comments = [];
          }

          array[i].isVote = this.isVote(array[i]._id);
          array[i].isBookmark = this.isBookmark(array[i]._id);
          parent.reply_comments.push(array[i]);
          parent.solutionCount = parent.reply_comments.filter(reply => reply.is_solution.value === true).length;
          parent.voteCount = parent.reply_comments.reduce((total, reply) => {
            if (reply.vote_count !== undefined) {
              return total + reply.vote_count;
            } else {
              return total;
            }
          }, 0);
          parent.replyCount = parent.reply_comments.length;
          const solutionComments = parent.reply_comments.filter(reply => reply.is_solution.value === true);
          if (!array[i].solutionLinks) {
            array[i].solutionLinks = [];
          }

          for (const solutionComment of solutionComments) {

            const commentLink = this.getCommentLink(solutionComment);
            if (solutionComment.is_solution.value !== undefined) {
              array[i].solutionLinks?.push(commentLink);

            }
            parent.solutionLinks = array[i].solutionLinks;
          }
          array[i].solutionCount = parent.solutionCount;
          array[i].voteCount = parent.voteCount;
          array[i].replyCount = parent.replyCount;
        }
      } else {
        array[i].isVote = this.isVote(array[i]._id);
        array[i].isBookmark = this.isBookmark(array[i]._id);
        tree.push(array[i]);
      }
    }
    return tree;
  }


  isVote(commentId: string) {
    return !!this.userVoteData.find((userVote: UserVote) => userVote.comment_id === commentId);
  }

  isBookmark(commentId: string) {
    return !!this.bookmarkData.find((bookmark: Bookmark) => bookmark.comment_id === commentId);
  }

  getCategoryList() {
    this.categorySvc.getCategoryList().subscribe({
      next: (categoryData: Category[]) => {
        this.categoryData = categoryData;
      }
    });
  }

  populateUniqueUserIds(): void {
    for (const comment of this.commentList) {
      this.uniqueUserIds.add(comment.user_id);

      if (comment.reply_comments) {
        const commentUserIdsSet = new Set<string>();

        for (const replyComment of comment.reply_comments) {
          if (replyComment.parent_id === comment._id) {
            commentUserIdsSet.add(replyComment.user_id);
          }
          this.uniqueUserIds.add(replyComment.user_id);
        }

        this.commentUserIds.set(comment._id, commentUserIdsSet);
      }
    }
  }

  getCategoryName(id: string): string {
    return this.categoryData.find((category: Category) => category._id === id)?.name || '';
  }

  getUserImageLink(id: string): string {
    return this.userData.find((user: User) => user._id === id)?.profile || '../../../assets/images/avatar.png';
  }

  getUserName(id: string): string {
    return this.userData.find((user: User) => user._id === id)?.display_name || '';
  }

  getUserNotiToken(id: string): string {
    return this.userData.find((user: User) => user._id === id)?.noti_token || '';
  }

  updateComment(comment: Comment): void {
    const data: UserModalData = {
      user: this.userData,
      parent_id: comment.parent_id,
      user_id: comment.user_id,
      question_id: comment.question_id,
      sort: comment.sort,
      type: 'update',
      comment_id: comment._id,
      description: comment.description,
      count: this.notiCount
    };
    this.callComment(data);
  }

  createAnswer(): void {
    if (this.userId != null) {
      const data: UserModalData = {
        user: this.userData,
        parent_id: "",
        user_id: this.userId,
        question_id: this.qaId,
        sort: this.commentList.length + 1,
        type: 'create',
        count: this.notiCount,
        question_detail_user_id: this.userId
      };
      this.callComment(data);
    }
  }

  createComment(comment: Comment): void {
    if (this.userId != null) {
      const data: UserModalData = {
        user: this.userData,
        parent_id: comment._id,
        user_id: this.userId,
        question_id: comment.question_id,
        sort: comment.reply_comments ? comment.reply_comments?.length + 1 : 1,
        type: 'create',
        count: this.notiCount,
        question_detail_user_id: comment.user_id
      };
      this.callComment(data);
    }
  }

  callComment(nzData: UserModalData) {
    const modal = this.modal.create({
      nzContent: CommentComponent,
      nzClosable: false,
      nzCentered: true,
      nzMaskClosable: false,
      nzStyle: { 'width': '80vw' },
      nzData: nzData,
    });
    modal.afterClose.subscribe((data) => {
      if (data === 'success') {
        this.getQaDetail();
      }
    });
  }

  shareComment(comment: Comment) {
    const commentLink = this.getCommentLink(comment);
    this.copyToClipboard(commentLink);
  }

  getCommentLink(comment: Comment): string {
    return `${this.baseUrl}/qa-detail/${comment.question_id}#${comment._id}`;
  }

  async copyToClipboard(text: string) {
    try {
      await navigator.clipboard.writeText(text);
      this.messageSvc.success(this.translate.instant('copyToClipboard'));
    } catch (error) {
      this.messageSvc.error(this.translate.instant('failToClipboard'));
    }
  }

  cancel(): void {
    this.messageSvc.info(this.translate.instant('cancelled'));
  }

  bookmark(comment: Comment): void {
    if (this.bookmarkClicked) {
      return;
    }
    this.bookmarkClicked = true;
    const bookmark = this.bookmarkData.find((bookmark: Bookmark) => bookmark.comment_id === comment._id) as Bookmark;
    if (comment.isBookmark) {
      this.bookmarkData = this.bookmarkData.filter((data) => {
        return data._id !== bookmark?._id;
      });
      comment.isBookmark = false;
      this.featureSvc.deleteBookmark(bookmark?._id || '').subscribe({
        next: () => {
          this.bookmarkClicked = false;
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error);
        }
      });
    } else {
      if (this.userId != null) {
        comment.isBookmark = true;
        const body: CreateBookmark = {
          user_id: this.userId,
          question_id: comment.question_id,
          comment_id: comment._id
        };
        this.featureSvc.createBookmark(body).subscribe({
          next: (data: Bookmark) => {
            this.bookmarkClicked = false;
            this.bookmarkData.push(data);
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.apiSvc.handleMessageError(error);
          }
        });
      }
    }
  }

  vote(comment: Comment): void {
    if (this.isClicked) {
      return;
    }
    this.isClicked = true;
    const userVote = this.userVoteData.find((userVote: UserVote) => userVote.comment_id === comment._id) as UserVote;
    if (comment.isVote) {
      this.questionDetail.vote_count = this.questionDetail.vote_count - 1;
      comment.vote_count = comment.vote_count - 1;

      comment.isVote = false;
      this.userVoteData = this.userVoteData.filter((data) => {
        return data._id !== userVote?._id;
      });
      this.featureSvc.deleteUserVote(userVote?._id || '').subscribe({
        next: () => {
          this.isClicked = false;
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error);
        }
      });
    } else {
      if (this.userId != null) {
        this.questionDetail.vote_count = this.questionDetail.vote_count + 1;
        comment.vote_count = comment.vote_count ? comment.vote_count + 1 : 1;
        comment.isVote = true;
        let tokenId = this.getUserNotiToken(comment.user_id)
        console.log(tokenId, "tokenId.......")
        const body: CreateUserVote = {
          user_id: this.userId,
          question_id: comment.question_id,
          comment_id: comment._id,
          noti_token: tokenId
        };
        this.featureSvc.createUserVote(body).subscribe({
          next: (data: UserVote) => {
            this.userVoteData.push(data);
            this.isClicked = false;
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.apiSvc.handleMessageError(error);
          }
        });
      }
    }
  }

  solution(comment: Comment): void {
    if (this.solutionClicked) {
      return;
    }
    this.solutionClicked = true;
    let tokenId = this.getUserNotiToken(comment.user_id)
    console.log(tokenId, "tokenId........")
    const body: RequestComment = {
      user_id: comment.user_id,
      parent_id: comment.parent_id,
      question_id: comment.question_id,
      sort: comment.sort,
      description: comment.description,
      is_solution: !comment.is_solution.value,
      noti_token: tokenId
    };

    this.commentSvc.updateComment(comment._id, body).subscribe({
      next: (data) => {
        comment.is_solution.value = !comment.is_solution.value;
        this.solutionClicked = false;
        this.badgeSvc.updateUserPoint(data.user_id, comment.is_solution.value).subscribe({
          next: () => {
            // if (comment.is_solution.value) {
            //   if (this.userId === data.user_id) {
            //     this.authSvc.noti.next(this.notiCount + 1)
            //   }
            // }
          }
        });
      },
      error: (error) => {
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
      }
    });
  }

  previewImg(src: string | undefined): void {
    const images = [{
      src: src ? src : '../../../assets/images/avatar.png'
    }];
    this.nzImageService.preview(images, { nzZoom: 1, nzRotate: 0 });
  }
}
