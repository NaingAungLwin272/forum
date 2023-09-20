import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { CategoryService } from 'src/app/repositories/category.service';
import { NzMessageService } from "ng-zorro-antd/message";
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-category-create',
  templateUrl: './category-create.component.html',
  styleUrls: ['./category-create.component.scss']
})
export class CategoryCreateComponent {
  @Input() categoryType: string = '';
  @Output() categoryCreated: EventEmitter<any> = new EventEmitter<any>();
  categoryCreateForm!: FormGroup;
  isLoading = false;

  constructor(
    private formBuilder: FormBuilder,
    private categorySvc: CategoryService,
    private messageSvc: NzMessageService,
    private modal: NzModalRef,
    private modalnew: NzModalService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.categoryCreateForm = this.formBuilder.group({
      name: ['', Validators.required]
    });
  }

  createCategory() {
    this.isLoading = true;
    this.loaderSvc.call();
    if (this.categoryCreateForm.valid) {
      const categoryData = {
        name: this.categoryCreateForm.value.name,
        type: this.categoryType === 'language' ? 1 : 2
      };
      this.categorySvc.createCategory(categoryData).subscribe({
        next: (categoryData) => {
          this.messageSvc.info(`category "${this.categoryCreateForm.value.name}" created successfully`);
          this.categoryCreated.emit(categoryData)
          this.isLoading = false;
          this.modal.destroy();
          this.loaderSvc.dismiss();

        },
        error: (error) => {
          if (error.status === 401) {
            this.loaderSvc.dissAll();
          }
          this.loaderSvc.dismiss();
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error)
          this.isLoading = false;
        }
      });
    }
  }

  showPopconfirm(): void {
    const confirmationMessage =
      this.categoryType === 'language'
        ? 'Are you sure you want to create this language?'
        : 'Are you sure you want to create this tag?';
    this.modalnew.confirm({
      nzTitle: confirmationMessage,
      nzOkText: 'OK',
      nzCancelText: 'Cancel',
      nzCentered: true,
      nzOnOk: () => {
        this.createCategory(),
          this.isLoading = false;
      },
      nzOnCancel: () => {
        this.modal.destroy();
      },
    });
  }

  public destroyModal(): void {
    this.modal.destroy();
  }
}
