import { Component, EventEmitter, Input, Output } from '@angular/core';
import { UserService } from 'src/app/repositories/user.service';
import { Department } from 'src/app/interfaces/department';
import { Team } from 'src/app/interfaces/team';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { User } from 'src/app/interfaces/user';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { LoaderService } from 'src/app/services/loader.service';
import { DisabledTimeFn, DisabledTimePartial } from 'ng-zorro-antd/date-picker';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-user-create',
  templateUrl: './user-create.component.html',
  styleUrls: ['./user-create.component.scss']
})
export class UserCreateComponent {
  @Input() teamData: Team[] = [];
  @Input() departmentData: Department[] = [];
  @Output() userCreated: EventEmitter<any> = new EventEmitter<any>()
  userData !: User;
  userForm !: FormGroup;
  passwordVisible = false;
  isConfirmLoading = false;
  today = new Date();
  options: string[] = [];

  constructor(
    private userSvc: UserService,
    private fb: FormBuilder,
    private modal: NzModalRef,
    private modalnew: NzModalService,
    private messageSvc: NzMessageService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) { }

  ngOnInit(): void {
    this.userForm = this.fb.group({
      display_name: ['', [Validators.required]],
      staff_id: ['', [Validators.required, Validators.pattern('^E[0-9]{5}$')]],
      name: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]],
      // password: ['', [Validators.required, Validators.minLength(6)]],
      role: ['', [Validators.required]],
      department_id: ['', [Validators.required]],
      team_id: ['', [Validators.required]],
      about_me: [''],
      // profile: [''],
      phone: ['', [Validators.pattern('^[0-9]{11}$')]],
      address: [''],
      dob: [null]
    });
    this.userForm.controls['department_id'].valueChanges.subscribe(data => {
      this.userForm.controls['team_id'].setValue('');
    })
  }

  cancelCreate() {
    this.modal.destroy();
  }

  createUser() {
    this.isConfirmLoading = true;
    this.loaderSvc.call();
    this.userSvc.createUser(this.userForm.value).subscribe({
      next: (userData: User) => {
        this.messageSvc.info(`user "${this.userForm.value.name}" created successfully`);
        this.userCreated.emit(userData)
        this.isConfirmLoading = false;
        this.modal.close();
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        if (error.status === 401) {
          this.loaderSvc.dissAll();
        }
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
        this.isConfirmLoading = false;
      }
    })
  }

  disabledDate = (current: Date): boolean => {
    return current.getTime() > Date.now();
  };

  showPopconfirm(): void {
    this.modalnew.confirm({
      nzTitle: 'Are you sure you want to create this user?',
      nzOkText: 'OK',
      nzCancelText: 'Cancel',
      nzCentered: true,
      nzOnOk: () => this.createUser(),
      nzOnCancel: () => {
        this.isConfirmLoading = false;
        this.modal.destroy();
      },
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
