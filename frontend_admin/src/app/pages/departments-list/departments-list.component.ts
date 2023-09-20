import { Component, ChangeDetectorRef, OnInit } from '@angular/core';
import { Department } from 'src/app/interfaces/department';
import { DepartmentService } from 'src/app/repositories/department.service';
import { NzModalService } from 'ng-zorro-antd/modal';
import { DepartmentCreateComponent } from 'src/app/components/department-create/department-create.component';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NzMessageService } from "ng-zorro-antd/message"
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-departments-list',
  templateUrl: './departments-list.component.html',
  styleUrls: ['./departments-list.component.scss']
})
export class DepartmentsListComponent implements OnInit {
  departmentEditForm!: FormGroup;
  clickedId!: string | '';
  editCache: { [key: string]: { edit: boolean; data: Department } } = {};
  departmentData: Department[] = []
  pageIndex!: number;
  pageSize!: number;
  departmentCount: number = 0;

  constructor(
    private departmentSvc: DepartmentService,
    public modal: NzModalService,
    private formBuilder: FormBuilder,
    private messageSvc: NzMessageService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.departmentEditForm = this.formBuilder.group({
      name: ['', Validators.required]
    });
  }

  ngOnInit(): void {
    this.pageIndex = 1;
    this.pageSize = 10;
    this.getDepartmentCount()
    this.getDepartmentLists(this.pageIndex, this.pageSize);
  }

  getDepartmentCount(): void {
    this.departmentSvc.getDepartmentCount().subscribe({
      next: (departmentCount) => {
        this.departmentCount = departmentCount.count;
      }, error: (error) => {
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  getDepartmentLists(pageIndex: number, pageSize: number) {
    this.loaderSvc.call();
    this.departmentSvc.getDepartmentList(pageIndex, pageSize).subscribe({
      next: (departmentListData: Department[]) => {
        this.departmentData = departmentListData
        this.loaderSvc.dissAll();
      },
      error: (error) => {
        this.loaderSvc.dissAll();
        this.departmentSvc.handleErrorType(error);
      }
    })
  }

  startEdit(id: string): void {
    this.clickedId = id;
    const selectedData = this.departmentData.find(data => data._id === id);
    this.departmentEditForm.patchValue({
      name: selectedData?.name
    })
  }

  cancelEdit(): void {
    this.clickedId = '';
  }

  saveEdit(id: string): void {
    if (this.departmentEditForm.valid) {
      const formData = this.departmentEditForm.value;
      this.loaderSvc.call();
      this.departmentSvc.updateDepartment(id, formData).subscribe({
        next: () => {
          this.messageSvc.info(`department "${this.departmentEditForm.value.name}" updated successfully`)
          this.getDepartmentLists(this.pageIndex, this.pageSize);
          this.clickedId = ""
          this.loaderSvc.dismiss();
        },
        error: (error) => {
          this.loaderSvc.dismiss();
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error);
        }
      })
    }
  }

  confirmDelete(id: string): void {
    this.loaderSvc.call();
    this.departmentSvc.deleteDepartment(id).subscribe({
      next: () => {
        this.messageSvc.info(`department deleted successfully`);
        this.getDepartmentCount();
        this.getDepartmentLists(this.pageIndex, this.pageSize);
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
      }
    })
  }

  cancel() {
    this.messageSvc.info("Cancelled")
  }

  createDepartment(): void {
    const modalRef = this.modal.create({
      nzContent: DepartmentCreateComponent,
      nzClosable: true,
      nzCentered: true,
      nzStyle: { 'width': '530px' },
      nzNoAnimation: true,
      nzMaskClosable: false,
    });
    if (this.departmentData) {
      modalRef?.componentInstance?.departmentCreated.subscribe((departmentData: Department) => {
        this.departmentData = [...this.departmentData, departmentData];
        this.getDepartmentCount();
      });
    }
  }

  pageIndexChanged(pageIndex: number): void {
    this.pageIndex = pageIndex;
    this.getDepartmentLists(pageIndex, this.pageSize);
  }
}
