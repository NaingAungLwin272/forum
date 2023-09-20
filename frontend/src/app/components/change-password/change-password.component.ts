import { Component, EventEmitter, Output } from '@angular/core';
import { AbstractControl, AbstractControlOptions, FormBuilder, FormGroup, UntypedFormControl, ValidatorFn, Validators } from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { AuthService } from 'src/app/repositories/auth.service';
import { ApiService } from 'src/app/services/api.service';
import { LoaderService } from 'src/app/services/loader.service';
import { TranslateService } from '@ngx-translate/core';


@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.scss']
})
export class ChangePasswordComponent {
  @Output() modalOpened: EventEmitter<void> = new EventEmitter<void>();
  changePasswordForm!: FormGroup;
  isConfirmLoading = false;
  passwordVisible = false;
  newPassVisible = false;
  confirmPassVisible = false;
  userId!: string;

  constructor(
    private formBuilder: FormBuilder,
    private modal: NzModalRef,
    private modalnew: NzModalService,
    private authSvc: AuthService,
    private apiSvc: ApiService,
    private messageSvc: NzMessageService,
    private loaderSvc: LoaderService,
    private translate: TranslateService,
  ) {
    this.userId = localStorage.getItem("login_user_id") || '';
    this.changePasswordForm = this.formBuilder.group({
      password: ['', [Validators.required, Validators.minLength(6)]],
      new_password: ['', [Validators.required, Validators.minLength(6)]],
      confirm_password: ['', [this.confirmValidator]]
    })
  }

  confirmValidator = (control: UntypedFormControl): { [s: string]: boolean } => {
    if (!control.value) {
      return { error: true, required: true };
    } else if (control.value !== this.changePasswordForm.controls['new_password'].value) {
      return { confirm: true, error: true };
    }
    return {};
  };

  cancel() {
    this.modal.close();
  }

  changePassword() {
    this.isConfirmLoading = true;
    this.loaderSvc.call();
    this.authSvc.changePassword(this.userId, this.changePasswordForm.value).subscribe({
      next: () => {
        this.messageSvc.success("Password Updated successfully")
        this.isConfirmLoading = false;
        this.loaderSvc.dismiss();
        this.modal.destroy();
      },
      error: (error) => {
        if (error.status === 401) {
          this.loaderSvc.dissAll();
        }
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error)
        this.apiSvc.handleMessageError(error)
        this.isConfirmLoading = false;
      }
    });
  }


  showChangePasswordPopUp() {
    this.modalnew.confirm({
      nzTitle: this.translate.instant('confirmPasswordChange'),
      nzOkText: this.translate.instant('ok'),
      nzCancelText: this.translate.instant('cancel'),
      nzCentered: true,
      nzOnOk: () => {
        this.changePassword()
      },

      nzOnCancel: () => {
        this.isConfirmLoading = false;
        this.modal.destroy();
      },
    });
  }
}
