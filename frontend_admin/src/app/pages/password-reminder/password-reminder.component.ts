import { Component, ViewEncapsulation } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

import { AuthService } from 'src/app/services/auth.service';
import { NzModalService } from 'ng-zorro-antd/modal';
import { Router } from '@angular/router';

@Component({
  selector: 'app-password-reminder',
  templateUrl: './password-reminder.component.html',
  styleUrls: ['./password-reminder.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class PasswordReminderComponent {
  mailForm: FormGroup;
  isLoadingTwo = false;

  constructor(
    private authSvc: AuthService,
    private router: Router,
    private modal: NzModalService
  ) {
    this.authSvc.clearAll();
    this.mailForm = new FormGroup({
      email: new FormControl('', [Validators.required, Validators.email]),
    });
  }

  passwordReminder(): void {
    this.isLoadingTwo = true;
    const body = {
      email: this.mailForm.value.email
    };

    const email = this.mailForm.value.email;

    this.authSvc.forgetPassowrd(email).then((data => {
      if (data === true) {
        this.isLoadingTwo = false;

        this.modal.success({
          nzTitle: 'Email Sent Successfully!',
          nzCentered: true,
          nzOnOk: () => {
            this.router.navigateByUrl('signin');
          }
        });
      }
    })
    );
  }
}
