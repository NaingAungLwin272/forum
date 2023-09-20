import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../services/api.service';
import { SigninUser } from '../interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class AuthService extends ApiService {
  clearAll() {
    throw new Error('Method not implemented.');
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
      { email, origin: "admin_page" }
    );
  }

  resetPassword(body: { email: string; token: string, password: string; }): Observable<any> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/auth/reset-password`,
      body
    );
  }
}
