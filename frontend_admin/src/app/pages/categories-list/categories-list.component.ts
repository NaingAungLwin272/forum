import { Component, OnInit } from '@angular/core';
import { Category } from 'src/app/interfaces/category';
import { CategoryService } from 'src/app/repositories/category.service';
import { NzMessageService } from "ng-zorro-antd/message"
import { NzModalService } from 'ng-zorro-antd/modal';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { CategoryCreateComponent } from 'src/app/components/category-create/category-create.component';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-categories-list',
  templateUrl: './categories-list.component.html',
  styleUrls: ['./categories-list.component.scss']
})
export class CategoriesListComponent implements OnInit {
  categoryData: Category[] = []
  languageData: Category[] = []
  tagData: Category[] = []
  categoryEditForm!: FormGroup;
  clickedId!: string | '';
  categoryType!: string;
  totalLanguages: number = 0;
  totalTags: number = 0;
  pageIndexLanguage!: number;
  pageIndexTag!: number;
  pageSize!: number;
  categoryCount: number = 0;
  languageCount: number = 0;
  tagCount: number = 0;

  constructor(
    private categorySvc: CategoryService,
    private messageSvc: NzMessageService,
    private formBuilder: FormBuilder,
    public modal: NzModalService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.categoryEditForm = this.formBuilder.group({
      name: ['', Validators.required]
    });
  }

  ngOnInit(): void {
    this.pageIndexLanguage = 1;
    this.pageIndexTag = 1;
    this.pageSize = 10;
    this.getCategoryCount()
    this.getCategoryByType(1, this.pageIndexLanguage, this.pageSize);
    this.getCategoryByType(2, this.pageIndexTag, this.pageSize)
  }

  getCategoryCount(): void {
    this.categorySvc.getCategoryCount().subscribe({
      next: (categoryCount) => {
        this.languageCount = categoryCount.language_count;
        this.tagCount = categoryCount.tag_count;
      }, error: (error) => {
        this.apiSvc.handleMessageError(error)
      }
    })
  }

  getCategoryByType(type?: number, pageIndex?: number, pageSize?: number) {
    this.loaderSvc.call();
    this.categorySvc.getCategoryByType(type, pageIndex, pageSize).subscribe({
      next: (categoryListData: Category[]) => {
        if (type == 1) {
          this.languageData = categoryListData;
        } else if (type == 2) {
          this.tagData = categoryListData
        }
        this.loaderSvc.dissAll();
      },
      error: (error) => {
        this.loaderSvc.dissAll();
        this.categorySvc.handleErrorType(error);
      }
    })
  }

  startEdit(id: string): void {
    this.clickedId = id;
    const selectedData = this.languageData.find(data => data._id === id);
    this.categoryEditForm.patchValue({
      name: selectedData?.name
    })
  }

  startEditTag(id: string): void {
    this.clickedId = id;
    const selectedData = this.tagData.find(data => data._id === id);
    this.categoryEditForm.patchValue({
      name: selectedData?.name
    })
  }

  confirmDelete(id: string, type: string): void {
    this.loaderSvc.call();
    this.categorySvc.deleteCategory(id).subscribe({
      next: () => {
        this.getCategoryCount()
        if (type == 'language') {
          this.getCategoryByType(1, this.pageIndexLanguage, this.pageSize);
        } else if (type == 'tag') {
          this.getCategoryByType(2, this.pageIndexTag, this.pageSize)
        }
        this.messageSvc.info(`category deleted successfully`);
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  saveEdit(id: string): void {
    this.loaderSvc.call();
    if (this.categoryEditForm.valid) {
      const formData = this.categoryEditForm.value;
      this.categorySvc.updateCategory(id, formData).subscribe({
        next: () => {
          this.messageSvc.info(`category "${this.categoryEditForm.value.name}" updated successfully`)
          this.getCategoryByType(1, this.pageIndexLanguage, this.pageSize);
          this.getCategoryByType(2, this.pageIndexTag, this.pageSize)
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

  pageIndexChanged(pageIndex: number, page: string): void {
    if (page == 'language') {
      this.pageIndexLanguage = pageIndex;
      this.getCategoryByType(1, this.pageIndexLanguage, this.pageSize);
    } else if (page == 'tag') {
      this.pageIndexTag = pageIndex;
      this.getCategoryByType(2, this.pageIndexTag, this.pageSize)
    }
  }

  cancelEdit(): void {
    this.clickedId = '';
  }

  cancel() {
    this.messageSvc.info("Cancel")
  }

  createCategory(categoryType: string) {
    const drawerRef = this.modal.create({
      nzContent: CategoryCreateComponent,
      nzClosable: true,
      nzCentered: true,
      nzStyle: { 'width': '530px' },
      nzNoAnimation: true,
      nzMaskClosable: false,
      nzComponentParams: {
        categoryType: categoryType
      }
    });
    if (this.categoryData) {
      drawerRef?.componentInstance?.categoryCreated.subscribe((categoryData) => {
        this.categoryType = categoryType
        if (this.categoryType === 'language') {
          this.languageData = [...this.languageData, categoryData];
        }
        if (this.categoryType === 'tag') {
          this.tagData = [...this.tagData, categoryData];
        }
        this.getCategoryCount();
      })
    }
  }
}
