import { Injectable } from '@angular/core';
import { CanActivate, Router, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from '../services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(
    private router: Router,
    private authSvc: AuthService
  ) {
  }
  canActivate(): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    const isLoggedIn = this.authSvc.isAuthenticated();
    if (isLoggedIn) {
      return true;
    }
    else {
      this.authSvc.clearAll();
      this.router.navigate(['/signin']);
      return false;
    }
  }

}
