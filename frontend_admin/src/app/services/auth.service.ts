import { BehaviorSubject } from 'rxjs/internal/BehaviorSubject';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';
import { reject } from 'lodash';
import { ApiService } from './api.service';
import { LoaderService } from './loader.service';
import { SigninUser } from '../interfaces/user';
import { AuthService as authSvc } from '../repositories/auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  public access_token = new BehaviorSubject<string | null>(null);
  public access_token$ = this.access_token.asObservable();

  public username = new BehaviorSubject<string>('');
  public username$ = this.username.asObservable();

  constructor(private authSvc: authSvc, private loaderService: LoaderService, private apiSvc: ApiService, private messageSvc: NzMessageService) {

    this.access_token.next(String(localStorage.getItem('access_token') || null));
  }

  isAuthenticated(): boolean {
    return localStorage.getItem('access_token') ? true : false;
  }

  signin(signin: SigninUser): Promise<{ is_authenticated: boolean, is_authorized: boolean; }> {
    return new Promise((resolve) => {
      this.authSvc.signin(signin).subscribe({
        next: (data) => {          
          if (data?.access_token && (data.user?.role === "manager" || data.user?.role === "bse" || data.user?.role === "leader")) {
            localStorage.setItem('access_token', data?.access_token);
            localStorage.setItem('login_user_id', data.user._id);

            this.access_token.next(data?.access_token);
            resolve({ is_authenticated: true, is_authorized: true });
          } else {
            this.loaderService.dismiss();
              resolve({ is_authenticated: true, is_authorized: false });
              this.messageSvc.error(
                "You are not authorized for this page."
              )
          }
        },
        error: (error) => {
          resolve({ is_authenticated: false, is_authorized: false });
          this.loaderService.dismiss();
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error);
        }
      });

    });
  }

  forgetPassowrd(email: string): Promise<any> {
    return new Promise((resolve) => {
      this.authSvc.forgetPassword(email).subscribe({
        next: (data) => {
          resolve(true)
        },
        error: (err) => {
          console.log("Login Error : ", err)
          this.apiSvc.handleMessageError(err);
        }
      })
    });
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
        }
      });
    });
  }

  clearAll(): void {
    localStorage.clear();
    this.access_token.next(null);
  }
}
