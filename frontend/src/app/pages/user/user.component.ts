import { Question } from 'src/app/interfaces/question';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NzUploadChangeParam, NzUploadComponent, NzUploadFile } from 'ng-zorro-antd/upload';
import { Observable, Observer } from 'rxjs';
import { Department } from 'src/app/interfaces/department';
import { DepartmentService } from 'src/app/repositories/department.service';
import { FeaturesService } from 'src/app/repositories/features.service';
import { QuestionService } from 'src/app/repositories/question.service';
import { Team } from 'src/app/interfaces/team';
import { TeamService } from 'src/app/repositories/team.service';
import { User } from 'src/app/interfaces/user';
import { UserService } from 'src/app/repositories/user.service';
import { UserSummary } from 'src/app/interfaces/user';
import { Vote } from 'src/app/interfaces/vote';
import { LoaderService } from '../../services/loader.service';
import { Category } from 'src/app/interfaces/category';
import { NotiService } from 'src/app/repositories/noti.service';
import { HttpClient } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { NzMessageService } from "ng-zorro-antd/message";
import { Noti, MapNoti } from 'src/app/interfaces/noti';
import { TranslateService } from '@ngx-translate/core';
import { UserBadge, Badge } from '../../interfaces/badge';
import { NzImageService } from 'ng-zorro-antd/image';
import { ApiService } from 'src/app/services/api.service';
import { AuthService } from 'src/app/services/auth.service';
import { DatePipe } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { ChangePasswordComponent } from 'src/app/components/change-password/change-password.component';
import { NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { MasterService } from 'src/app/services/master.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.scss']
})
export class UserComponent implements OnInit {

  @ViewChild('uploadComponent') uploadComponent!: NzUploadComponent;
  selectedTabIndex = 0;
  userInfo?: User;
  team?: Team;
  bookmark!: Vote[];
  vote!: Vote[];
  noti!: Noti[];
  commentData: Vote[] = [];
  answerData!: Vote[];
  commentDataSolved!: Vote[];
  department?: Department;
  questionData!: Question[];
  mentionData!: Vote[];
  userSummary = {} as UserSummary;
  editProfileForm: FormGroup;
  isLoading = false;
  currentPage = 1;
  imageUrl: string | undefined | null;
  categoryData: Category[] = [];
  fileList: any;
  mail_subscribe!: boolean;
  mappedNoti!: MapNoti[];
  roleInt!: number;
  notiCount = 0;
  notificationsMarkedAsRead = false;
  lastpost = '';
  loggedUserId = localStorage.getItem('login_user_id');
  userBadgeData: UserBadge[] = [];
  badgeData: Badge[] = [];
  userProfile!: string;
  userId!: string;
  commonUserId!: string;
  user!: User;
  darkTheme!: boolean;
  options: string[] = [];

  constructor(
    public loaderService: LoaderService,
    public teamService: TeamService,
    public departmentService: DepartmentService,
    public featureService: FeaturesService,
    private formBuilder: FormBuilder,
    private questionSvc: QuestionService,
    private userSvc: UserService,
    private messageSvc: NzMessageService,
    private notiSvc: NotiService,
    private http: HttpClient,
    private translate: TranslateService,
    private nzImageService: NzImageService,
    private apiSvc: ApiService,
    private authSvc: AuthService,
    private datePipe: DatePipe,
    private route: ActivatedRoute,
    private modalSvc: NzModalService,
    private authService: AuthService,
    private router: Router,
    private masterSvc: MasterService
  ) {
    this.editProfileForm = this.formBuilder.group({
      display_name: ['', [Validators.required]],
      staff_id: ['', [Validators.required, Validators.pattern('^E[0-9]{5}$')]],
      name: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]],
      role: [''],
      about_me: [''],
      profile: [''],
      phone: ['', [Validators.pattern('^[0-9]{11}$')]],
      address: [''],
      mail_subscribe: [false],
      dob: [null]
    });
  }

  ngOnInit(): void {
    this.getUserInfo();
    this.darkTheme = localStorage.getItem('theme') === 'dark';
    this.authService.noti$.subscribe(async (data: number) => {
      this.notiCount = data;
    });
    this.masterSvc.categoryData$.subscribe((categoryData: Category[]) => {
      this.categoryData = categoryData;
    });
  }

  getUserInfo() {
    this.route.paramMap.subscribe(params => {
      this.userId = params.get('id') || '';
      this.loaderService.call();
      this.userSvc.getUserProfile(this.userId).subscribe({
        next: async (user: User) => {
          this.userInfo = user;
          if (this.userInfo) {
            await this.getTeam(this.userInfo.team_id);
            await this.getDepartment(this.userInfo.department_id);
            this.getUserSummary(this.userInfo._id);
            this.getQuestionsByUserId(this.userInfo._id, this.currentPage, 10);
            this.getAnswersByUserId(this.userInfo._id, this.currentPage, 10);
            this.getBookmarkByUserId(this.userInfo._id, this.currentPage, 10);
            this.getVoteByUserId(this.userInfo._id, this.currentPage, 10);
            this.getCommentByUserIdSolved(this.userInfo._id, this.currentPage, 10);
            this.getMentionByUserId(this.userInfo._id, this.currentPage, 10);
            this.getNotiByUserId(this.userInfo._id, this.currentPage, 10);
            this.getBadges();
            this.getBadgesByUserId(this.userInfo._id);

            const seconds = this.userInfo.dob.seconds;
            if (!isNaN(seconds)) {
              const milliseconds = seconds * 1000;
              const date = new Date(milliseconds);
              this.editProfileForm.patchValue({
                dob: date,
              });
            }
            this.editProfileForm.patchValue({
              staff_id: this.userInfo.staff_id,
              team_id: this.userInfo.team_id,
              email: this.userInfo.email,
              name: this.userInfo.name,
              display_name: this.userInfo.display_name,
              phone: this.userInfo.phone,
              role: this.userInfo.role,
              department_id: this.userInfo.department_id,
              about_me: this.userInfo.about_me,
              address: this.userInfo.address,
              profile: this.userInfo.profile,
              mail_subscribe: !!this.userInfo.mail_subscribe
            });
            this.imageUrl = this.userInfo.profile;
            this.editProfileForm.get('staff_id')?.disable();
            this.editProfileForm.get('role')?.disable();
            this.loaderService.dismiss();
          }
        },
        error: (error) => {
          this.loaderService.dismiss();
          this.apiSvc.handleErrorType(error)
          this.apiSvc.handleMessageError(error)
          this.router.navigate(['/'])
        }
      });
    })
  }

  navigateToTabs(index: number) {
    this.selectedTabIndex = index;
    if (this.selectedTabIndex == 8 && this.userId != null) {
      this.MarkAllNotiAsRead(this.userId)
    }
  }

  MarkAllNotiAsRead(id: string) {
    this.notiSvc.MarkAllNotiAsRead(id).subscribe({
      next: () => {
        this.authSvc.noti.next(0);
        this.loaderService.dismiss();
      },
      error: () => {
        this.loaderService.dismiss();
      }
    })
  }

  async getTeam(id: string) {
    return new Promise(resovle => {
      this.teamService.getTeam(id).subscribe({
        next: (team: Team) => {
          this.team = team;
          resovle(true);
        },
        error: () => {
          resovle(false);
          this.loaderService.dismiss();
        }
      });
    })
  }

  async getDepartment(id: string) {
    return new Promise(resovle => {
      this.departmentService.getDepartment(id).subscribe({
        next: (department: Department) => {
          this.department = department;
          resovle(true);
        },
        error: () => {
          resovle(false);
          this.loaderService.dismiss();
        }
      });
    });
  }

  getQuestionsByUserId(userId: string, page: number, limit: number) {
    this.questionSvc.getQuestionsWithUserId(userId, page, limit).subscribe({
      next: (questionData: Question[]) => {
        this.questionData = questionData;
        this.loaderService.dismiss();
      },
      error: (error) => {
        this.loaderService.dismiss();
        this.apiSvc.handleMessageError(error)
      }
    })
  }

  getUserSummary(userId: string) {
    this.userSvc.getUserSummary(userId).subscribe({
      next: (userSummary: UserSummary) => {
        this.userSummary = userSummary;
        this.loaderService.dismiss();
      },
      error: () => {
        this.loaderService.dismiss();
      }
    })
  }

  getBookmarkByUserId(id: string, page: number, limit: number) {
    this.featureService.getBookmarkByUserId(id, page, limit).subscribe({
      next: (bookmark: Vote[]) => {
        this.bookmark = bookmark;
        this.loaderService.dismiss();
      },
      error: () => {
        this.loaderService.dismiss();
      }
    });
  }

  getVoteByUserId(id: string, page: number, limit: number) {
    this.featureService.getVotesByUserId(id, page, limit).subscribe({
      next: (vote: Vote[]) => {
        this.vote = vote;
        this.loaderService.dismiss();
      },
      error: () => {
        this.loaderService.dismiss();
      }
    });
  }

  getAnswersByUserId(id: string, page: number, limit: number) {
    this.featureService.getAnswersByUserId(id, page, limit).subscribe({
      next: (commentData: any[]) => {
        this.answerData = commentData;
        this.loaderService.dismiss();
      },
      error: () => {
        this.loaderService.dismiss();
      }
    })
  }

  getCommentByUserIdSolved(id: string, page: number, limit: number) {
    this.featureService.getCommentByUserIdSolved(id, page, limit).subscribe({
      next: (commentDataSolved: Vote[]) => {
        this.commentDataSolved = commentDataSolved;
        this.loaderService.dismiss();
      },
      error: () => {
        this.loaderService.dismiss();
      }
    })
  }

  getMentionByUserId(id: string, page: number, limit: number) {
    this.featureService.getMentionByUserId(id, page, limit).subscribe({
      next: (mentionData: Vote[]) => {
        this.mentionData = mentionData;
        this.loaderService.dismiss();
      },
      error: () => {
        this.loaderService.dismiss();
      }
    })
  }

  getBadges() {
    this.featureService.getBadgeList().subscribe({
      next: (badgeData: Badge[]) => {
        this.badgeData = badgeData;
        this.loaderService.dismiss();
      },
      error: (error) => {
        this.apiSvc.handleMessageError(error);
        this.loaderService.dismiss();
      }
    })
  }

  getBadgesByUserId(id: string) {
    this.featureService.getBadgesByUserId(id).subscribe({
      next: (userBadgeData: UserBadge[]) => {
        this.userBadgeData = userBadgeData || [];
        this.loaderService.dismiss();
      },
      error: (error) => {
        this.apiSvc.handleMessageError(error);
        this.loaderService.dismiss();
      }
    })
  }

  isUserBadge(id: string) {
    return this.userBadgeData?.find(userBadge => userBadge.badge_id === id)?._id;
  }

  formatDateTime(created_at: { seconds: number; nanos: number } | string | null): string {
    if (typeof created_at === 'string') {
      const date = new Date(created_at);
      return this.datePipe.transform(date, 'yyyy-MM-dd HH:mm:ss') ?? '';
    } else if (created_at && typeof created_at === 'object') {
      const timestampInSeconds = created_at.seconds;
      const timestampInMilliseconds = timestampInSeconds * 1000 + created_at.nanos / 1000000;
      const date = new Date(timestampInMilliseconds);
      return this.datePipe.transform(date, 'yyyy-MM-dd HH:mm:ss') ?? '';
    } else {
      return '';
    }
  }

  getNotiByUserId(userId: string, page: number, limit: number) {
    this.notiSvc.getNotiByUserId(userId, page, limit).subscribe({
      next: (noti: Noti[]) => {
        this.noti = noti;
        this.mappedNoti = this.noti.map(notification => ({
          title: notification.name,
          description: notification.description,
          type: notification.type,
          link: notification.link,
          created_at: this.formatDateTime(notification.created_at)
        }));
        this.loaderService.dismiss();
      },
      error: (error) => {
        this.loaderService.dismiss();
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  getRibbonText(type: number): string {
    switch (type) {
      case 1:
        return this.translate.instant('levelUp');
      case 2:
        return this.translate.instant('replyNoti');
      case 3:
        return this.translate.instant('mentionNoti');
      case 4:
        return this.translate.instant('solveNoti');
      case 5:
        return this.translate.instant('voteNoti');
      case 6:
        return this.translate.instant('systemNoti');
      default:
        return this.translate.instant('Unknown');
    }
  }

  getRibbonColor(type: number): string {
    switch (type) {
      case 1:
        return 'pink';
      case 2:
        return 'cyan';
      case 3:
        return 'green';
      case 4:
        return 'purple';
      case 5:
        return 'volcano';
      case 6:
        return 'geekblue';
      default:
        return 'yellow';
    }
  }

  beforeUpload = (file: NzUploadFile): Observable<boolean> =>
    new Observable((observer: Observer<boolean>) => {
      const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png' || file.type == 'image/jpg';
      if (!isJpgOrPng) {
        this.messageSvc.error(this.translate.instant('imageOnly'));
        observer.complete();
        return;
      }
      const isLt2M = file.size! / 1024 / 1024 < 2;
      if (!isLt2M) {
        this.messageSvc.error(this.translate.instant('twoMb'));
        observer.complete();
        return;
      }
      observer.next(isJpgOrPng && isLt2M);
      observer.complete();
    });

  private getBase64(img: File, callback: (img: string) => void): void {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result!.toString()));
    reader.readAsDataURL(img);
  }

  handleChange(event: NzUploadChangeParam): void {
    this.fileList = [...event.fileList];
    if (this.fileList.length > 0 && this.fileList[0].originFileObj instanceof File) {
      this.getBase64(this.fileList[this.fileList.length - 1].originFileObj, (img: string) => {
        this.isLoading = false;
        this.imageUrl = img;
      });
    } else {
      this.imageUrl = null;
    }
  }

  uploadImage(file: File) {
    const uploadData = new FormData();
    uploadData.append('file', file);
    uploadData.append('upload_preset', environment.preset);
    uploadData.append('folder', 'mtm_community_profile');
    return this.http.post(`https://api.cloudinary.com/v1_1/${environment.cloudName}/upload?upload_preset=${environment.preset}`, uploadData);
  }

  editProfile() {
    if (this.editProfileForm.valid) {
      this.isLoading = true;
      const role = this.editProfileForm.get('role')?.value;
      switch (role) {
        case "manager":
          this.roleInt = 1
          break;
        case "bse":
          this.roleInt = 2
          break;
        case "leader":
          this.roleInt = 3
          break;
        case "sub leader":
          this.roleInt = 4;
          break;
        case "senior":
          this.roleInt = 5;
          break;
        case "junior":
          this.roleInt = 6;
          break
        default:
          this.roleInt = 6;
      }
      const formData = { ...this.editProfileForm.value, role: this.roleInt };
      if (this.fileList && this.fileList.length > 0 && this.fileList[0].originFileObj instanceof File) {
        const file = this.fileList[this.fileList.length - 1].originFileObj;
        this.loaderService.call();
        this.uploadImage(file).subscribe({
          next: (response: any) => {
            this.imageUrl = response.url;
            formData.profile = this.imageUrl;
            this.userSvc.editProfile(this.userInfo?._id, formData).subscribe({
              next: (data: User) => {
                this.getUserInfo();
                this.authSvc.user_profile.next(data.profile);
                this.messageSvc.success(this.translate.instant('profileUpdate'));
                this.isLoading = false;
                this.loaderService.dissAll();
              },
              error: (error) => {
                this.loaderService.dissAll();
                this.apiSvc.handleMessageError(error);
                this.apiSvc.handleErrorType(error);
                this.isLoading = false;
              }
            });
          },
          error: (error) => {
            this.loaderService.dissAll()
            this.apiSvc.handleMessageError(error);
            this.apiSvc.handleErrorType(error);
          }
        });

        this.fileList.length = 0;
      } else {
        this.loaderService.call();
        this.userSvc.editProfile(this.userInfo?._id, formData).subscribe({
          next: () => {
            this.getUserInfo();
            this.messageSvc.success(this.translate.instant('profileUpdate'));
            this.isLoading = false;
            this.loaderService.dissAll();
          },
          error: (error) => {
            this.isLoading = false;
            this.apiSvc.handleMessageError(error);
            this.apiSvc.handleErrorType(error);
            this.loaderService.dissAll();
          }
        });
      }
    } else {
      this.messageSvc.error(this.translate.instant('invalidFormData'));
      this.loaderService.dissAll();
      this.isLoading = false;
    }
  }

  loadMoreNoti() {
    this.loaderService.call()
    this.currentPage++;
    const limit = 10;
    if (this.userId !== null) {
      this.notiSvc.getNotiByUserId(this.userId, this.currentPage, limit).subscribe({
        next: (data: Noti[]) => {
          this.noti = [...this.noti, ...data]
          this.mappedNoti = this.noti.map(notification => ({
            title: notification.name,
            description: notification.description,
            type: notification.type,
            link: notification.link,
            created_at: this.formatDateTime(notification.created_at)
          }));
          this.loaderService.dismiss();
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error);
          this.loaderService.dismiss();
        }
      })
    }
  }

  navigateToLink(link: string) {
    console.log(link, "link......")
    window.open(link);
  }

  previewImg(src: string | undefined): void {
    const images = [{
      src: src ? src : '../../../assets/images/avatar.png'
    }];
    this.nzImageService.preview(images, { nzZoom: 1, nzRotate: 0 });
  }

  disabledDate = (current: Date): boolean => {
    const oneHundredYearsAgo = new Date();
    oneHundredYearsAgo.setFullYear(new Date().getFullYear() - 100);
    return current.getTime() > Date.now() || current.getTime() < oneHundredYearsAgo.getTime();
  };

  changePasswordModal() {
    this.modalSvc.create({
      nzContent: ChangePasswordComponent,
      nzFooter: null,
      nzClosable: true,
      nzCentered: true,
      nzStyle: { 'width': '40vw' },
      nzNoAnimation: false,
      nzMaskClosable: false,
    });
  }

  onInput(e: Event): void {
    const value = (e.target as HTMLInputElement).value;
    if (!value || value.indexOf('@') >= 0) {
      this.options = [];
    } else {
      this.options = ['gmail.com', 'metateammyanmar.com'].map(domain => `${value}@${domain}`);
    }
  }
}



