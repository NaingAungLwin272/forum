import { Component, ViewEncapsulation } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AuthService } from 'src/app/services/auth.service';
import { LoaderService } from 'src/app/services/loader.service';
import { NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class SigninComponent {
  signinForm: FormGroup;
  passwordVisible = false;
  options: string[] = [];
  constructor(
    private authSvc: AuthService,
    private loaderSvc: LoaderService,
    private router: Router,
    private modal: NzModalService
  ) {
    this.authSvc.clearAll();
    this.loaderSvc.dismiss();
    this.signinForm = new FormGroup({
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required]),
      is_remember_me: new FormControl(false)
    });
  }

  signin(): void {
    this.loaderSvc.call();
    this.authSvc.signin(this.signinForm.value).then((value: { is_authenticated: boolean, is_authorized: boolean }) => {
      if (value.is_authenticated === true && value.is_authorized === true) {
        this.loaderSvc.dismiss();
        this.router.navigate(['/']);
      }
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
