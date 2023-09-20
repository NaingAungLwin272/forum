import { Component, OnInit } from '@angular/core';

import { AuthService } from './services/auth.service';
import { ThemeService } from './services/theme.service';
import { TranslateService } from '@ngx-translate/core';
import { NotiService } from './repositories/noti.service';
import { UserService } from './repositories/user.service';
import { isEmpty } from 'lodash';
import { User } from './interfaces/user';
import { ApiService } from './services/api.service';
import { MasterService } from 'src/app/services/master.service';
import { getMessaging, getToken, onMessage } from 'firebase/messaging';
import { environment } from 'src/environments/environment.development';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  isLogin = false;
  currentTheme!: boolean;
  profile!: string;
  currentLanguage!: string;
  language!: string;
  notiCount = 0;
  userProfile!: string | null;
  userId!: string;
  userData!: User;
  notiToken!: string;

  constructor(
    private themeService: ThemeService,
    private authService: AuthService,
    private notiService: NotiService,
    private translate: TranslateService,
    private userSvc: UserService,
    private apiSvc: ApiService,
    private masterSvc: MasterService
  ) {
  }

  ngOnInit(): void {
    this.currentTheme = localStorage.getItem('theme') === 'dark' ? false : true;
    this.authService.access_token$.subscribe(async (access_token: string | null) => {
      await this.masterSvc.getCategoryList();
      await this.masterSvc.getUserList();
      this.userId = localStorage.getItem("login_user_id") || '';
      this.isLogin = !isEmpty(access_token) && access_token !== 'null' ? true : false;
      if (this.userId) {
        this.getNotiCount(this.userId)
        this.requestPermission();
      }
      this.currentLanguage = localStorage.getItem('language') || 'en';
      if (this.currentLanguage) {
        this.translate.use(this.currentLanguage);
      }
    });
    this.authService.user_profile$.subscribe(async (data: string | null) => {
      this.userProfile = data;
    });
    this.authService.noti$.subscribe(async (data) => {
      this.notiCount = data;
    })
  }

  // for realtime noti
  requestPermission() {
    const messaging = getMessaging();
    getToken(messaging, { vapidKey: environment.firebase.vapidKey }).then(
      (currentToken) => {
        if (currentToken) {
          this.getUser(this.userId, currentToken);
        } else {
          console.log("we have a problem");
        }
      }
    )
    onMessage(messaging, (payload) => {
      if (payload) {
        this.notiCount += 1
        this.authService.noti.next(this.notiCount)
      }
    });
  }

  getUser(userId: string, notiToken: string) {
    this.userSvc.getUser(userId).subscribe({
      next: (userData: User) => {
        this.authService.user_profile.next(userData.profile)
        this.userData = userData;
        const role = this.validateUserRole(userData.role)
        if (this.userData.noti_token === undefined) {
          const data = {
            role: role,
            noti_token: notiToken
          }
          this.userSvc.editProfile(this.userId, data).subscribe({
            next: (data) => {
              console.log(data, "data.....")
            }
          })
        }
      },
      error: (error) => {
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  validateUserRole(role: string): number {
    let roleNumber = 0;
    switch (role) {
      case 'manager':
        roleNumber = 1;
        break;
      case 'bse':
        roleNumber = 2;
        break;
      case 'leader':
        roleNumber = 3;
        break;
      case 'sub leader':
        roleNumber = 4;
        break;
      case 'senior':
        roleNumber = 5;
        break;
      case 'junior':
        roleNumber = 6;
        break;
    }
    return roleNumber
  }

  switchLanguage(language: string): void {
    this.translate.use(language);
    this.currentLanguage = language;
    localStorage.setItem('language', language);
  }

  getNotiCount(id: string) {
    this.notiService.getNotiCount(id).subscribe({
      next: (notiCount: { count: number }) => {
        this.authService.noti.next(notiCount.count);
        console.log(notiCount, "notiCount....")
        if (notiCount.count > 0) {
          
          this.notiCount = notiCount.count;
        } else {
          this.notiCount = 0;
        }
      }
    })
  }

  changeTheme(): void {
    this.themeService.toggleTheme().then();
  }
}
