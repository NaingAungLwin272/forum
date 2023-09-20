import { AbstractControl, FormControl, FormGroup, ValidatorFn, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Component } from '@angular/core';
import { AuthService } from 'src/app/services/auth.service';
import { NzModalService } from 'ng-zorro-antd/modal';

@Component({
  selector: 'app-password-reset',
  templateUrl: './password-reset.component.html',
  styleUrls: ['./password-reset.component.scss']
})
export class PasswordResetComponent {
  resetPasswordForm: FormGroup;
  isLoadingTwo = false;
  confirmPassVisible = false;
  passwordVisible = false;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private authSvc: AuthService,
    private modal: NzModalService
  ) {
    this.authSvc.clearAll();
    this.resetPasswordForm = new FormGroup({
      password: new FormControl('', [Validators.required, Validators.minLength(6)]),
      retype_password: new FormControl('', [Validators.required]),
    }, { validators: this.passwordMatchValidator });
  }

  // Custom validator to check if passwords match
  passwordMatchValidator: ValidatorFn = (control: AbstractControl): { [key: string]: any; } | null => {
    const passwordControl = control.get('password');
    const retypePasswordControl = control.get('retype_password');

    if (passwordControl?.value !== retypePasswordControl?.value) {
      retypePasswordControl?.setErrors({ passwordMismatch: true });
      return { passwordMismatch: true };
    } else {
      retypePasswordControl?.setErrors(null);
      return null;
    }
  };

  passwordReminder(): void {
    this.isLoadingTwo = true;
    const body = {
      password: this.resetPasswordForm.value.password,
      token: this.route.snapshot.queryParams['token'],
      email: this.route.snapshot.queryParams['email']
    };

    this.authSvc.resetPassword(body).then((data) => {
      if (data == true) {
        this.isLoadingTwo = false;
        this.modal.success({
          nzTitle: 'Password Changed Successfully!',
          nzCentered: true,
          nzOnOk: () => {
            this.router.navigateByUrl('signin');
          }
        });
      }
    });
  }
}
