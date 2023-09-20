import { Component, EventEmitter, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { DepartmentService } from 'src/app/repositories/department.service';
import { NzMessageService } from "ng-zorro-antd/message"
import { Department } from 'src/app/interfaces/department';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-department-create',
  templateUrl: './department-create.component.html',
  styleUrls: ['./department-create.component.scss']
})
export class DepartmentCreateComponent {
  @Output() departmentCreated: EventEmitter<any> = new EventEmitter<any>();
  departmentCreateForm!: FormGroup;
  isConfirmLoading = false;
  constructor(
    private modal: NzModalRef,
    private modalnew: NzModalService,
    private formBuilder: FormBuilder,
    private DepartmentSvc: DepartmentService,
    private messageSvc: NzMessageService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.departmentCreateForm = this.formBuilder.group({
      name: ['', Validators.required]
    });
  }

  public destroyModal(): void {
    this.modal.destroy();
  }

  createDepartment() {
    this.isConfirmLoading = true;
    this.loaderSvc.call();
    if (this.departmentCreateForm.valid) {
      const departmentData = this.departmentCreateForm.value;
      this.DepartmentSvc.createDepartment(departmentData).subscribe({
        next: (departmentData: Department) => {
          this.messageSvc.info(`Department "${this.departmentCreateForm.value.name}" Created successfully`);
          this.departmentCreated.emit(departmentData);
          this.isConfirmLoading = false;
          this.modal.destroy();
          this.loaderSvc.dismiss();
        },
        error: (error) => {
          if (error.status === 401) {
            this.loaderSvc.dissAll();
          }
          this.loaderSvc.dismiss();
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error);
        }
      })
    }
  }

  showPopconfirm(): void {
    this.modalnew.confirm({
      nzTitle: 'Are you sure you want to create this department?',
      nzOkText: 'OK',
      nzCancelText: 'Cancel',
      nzCentered: true,
      nzOnOk: () => {
        this.createDepartment(),
          this.isConfirmLoading = false;
      },
      nzOnCancel: () => {
        this.modal.destroy();
      },
    });
  }
}
