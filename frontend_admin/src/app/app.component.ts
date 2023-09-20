import { Component, OnInit } from '@angular/core';
import { Router, Event, NavigationStart, NavigationEnd } from '@angular/router';
import { isEmpty } from 'lodash';
import { AuthService } from './services/auth.service';
import { UserService } from './repositories/user.service';
import { User } from './interfaces/user';
import { ApiService } from './services/api.service';
import { LoaderService } from './services/loader.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  isCollapsed = true;
  isLogin = true;
  clientName!: string;
  currentRoute!: string;
  userData!: User;
  uId!: string;
  userName!: string;

  constructor(
    private authSvc: AuthService,
    private router: Router,
    private userSvc: UserService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.router.events.subscribe((event: Event) => {
      if (event instanceof NavigationStart) {
        this.isCollapsed = true;
      }
    });
  }

  ngOnInit(): void {
    this.router.events.subscribe(
      (event) => {
        if (event instanceof NavigationEnd) {
          this.currentRoute = this.router.url;
        }
      }
    );
    this.authSvc.access_token$.subscribe(async (access_token: string | null) => {
      this.isLogin = !isEmpty(access_token) && access_token !== 'null' ? true : false;
      if (!this.isLogin) {
        this.isCollapsed = true;
      }
      else {
        const uId = String(localStorage.getItem('login_user_id'));
        await this.getUser(uId);
      }
    });
    this.authSvc.username$.subscribe(async (data: string) => {
      this.userName = data;
    });
  }
  getUser(userId: string) {
    this.userSvc.getUser(userId).subscribe({
      next: (userData: User) => {
        this.loaderSvc.dismiss();
        this.authSvc.username.next(userData.name);
        this.userData = userData;
      },
      error: (error) => {
        this.apiSvc.handleMessageError(error);
      }
    })
  }
}
