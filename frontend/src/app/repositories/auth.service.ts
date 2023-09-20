import { SigninUser, User } from '../interfaces/user';

import { ApiService } from '../services/api.service';
import { Injectable } from '@angular/core';
import { Login } from '../interfaces/auth';
import { Observable } from 'rxjs';
import { PasswordChange } from '../interfaces/chagePassword';

@Injectable({
  providedIn: 'root'
})
export class AuthService extends ApiService {

  getUserProfile(): Observable<User> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/auth/local`,
      {}
    );
  }

  postLogin(formData: Login): Observable<any> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/auth/login`,
      formData
    );
  }

  signin(signin: SigninUser): Observable<any> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/auth/login`,
      signin
    );
  }

  forgetPassword(email: string): Observable<any> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/auth/forget-password`,
      { email, origin: "test" }
    );
  }

  resetPassword(body: { email: string; token: string, password: string; }): Observable<any> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/auth/reset-password`,
      body
    );
  }

  changePassword(userId: string, body: PasswordChange): Observable<any> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/user/change-password`,
      { user_id: userId, ...body }
    )
  }
}
