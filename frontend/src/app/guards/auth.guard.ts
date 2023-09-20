import { Injectable } from '@angular/core';
import { CanActivate, Router, UrlTree, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
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
  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    if (route.routeConfig?.path === 'user/:id') {
      const access_token = route.queryParams?.['access_token'];
      const login_user_id = route.queryParams?.['login_user_id'];
      if (access_token && login_user_id) {
        localStorage.setItem('access_token', access_token);
        localStorage.setItem('login_user_id', login_user_id);
        this.authSvc.access_token.next(access_token);
        return true;
      }
    }
    const isLoggedIn = this.authSvc.isAuthenticated();
    if (isLoggedIn) {
      return true;
    }
    else {
      this.authSvc.clearAll();
      this.router.navigate(['/signin'], { queryParams: { redirectUrl: state.url } });
      return false;
    }
  }

}
