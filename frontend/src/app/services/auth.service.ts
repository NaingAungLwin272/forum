import { BehaviorSubject } from 'rxjs/internal/BehaviorSubject';
import { Injectable } from '@angular/core';
import { LoaderService } from './loader.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { SigninUser } from '../interfaces/user';
import { AuthService as authSvc } from '../repositories/auth.service';
import { ApiService } from './api.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  public access_token = new BehaviorSubject<string | null>(null);
  public access_token$ = this.access_token.asObservable();

  public user_profile = new BehaviorSubject<string | null>(null);
  public user_profile$ = this.user_profile.asObservable();

  public noti = new BehaviorSubject<number>(0);
  public noti$ = this.noti.asObservable();

  constructor(private authSvc: authSvc, private loaderService: LoaderService, private messageSvc: NzMessageService, private apiSvc: ApiService) {
    this.access_token.next(String(localStorage.getItem('access_token') || null));
  }

  isAuthenticated(): boolean {
    return localStorage.getItem('access_token') ? true : false;
  }

  signin(signin: SigninUser): Promise<{ is_authenticated: boolean; }> {

    return new Promise((resolve) => {
      this.authSvc.signin(signin).subscribe({
        next: (data) => {
          localStorage.setItem('language', 'en');
          if (data?.access_token) {
            localStorage.setItem('access_token', data?.access_token);
            localStorage.setItem('login_user_id', data.user._id)
            this.access_token.next(data?.access_token);
            resolve({ is_authenticated: true })
          } else {
            resolve({ is_authenticated: false })
          }
        },
        error: (err) => {
          resolve({ is_authenticated: false })
          console.log("Login Error : ", err)
          this.loaderService.dismiss();
          this.apiSvc.handleErrorType(err);
          this.apiSvc.handleMessageError(err);
        }
      })

    });
  }

  forgetPassword(email: string): Promise<any> {

    return new Promise((resolve) => {
      this.authSvc.forgetPassword(email).subscribe({
        next: (data) => {
          console.log(data, "response password...")
          resolve(true)
        },
        error: (err) => {
          this.apiSvc.handleMessageError(err);
        }
      })
    })
  }

  resetPassword(body: { email: string; token: string; password: string }): Promise<any> {
    return new Promise((resolve) => {
      this.authSvc.resetPassword(body).subscribe({
        next: (data) => {
          if (data.is_success == true) {
            resolve(true);
          } else {
            resolve(false);
          }
        },
        error: (err) => {
          this.apiSvc.handleMessageError(err);
          // resolve(err);
        }
      });
    });
  }

  clearAll(): void {
    const theme = localStorage.getItem("theme") || '';
    localStorage.clear();
    localStorage.setItem("theme", theme);
    this.access_token.next(null);
  }
}
