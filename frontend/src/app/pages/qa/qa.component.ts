import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NzDrawerService } from 'ng-zorro-antd/drawer';
import { TranslateService } from '@ngx-translate/core';
import { Category } from '../../interfaces/category';
import { CategoryService } from '../../repositories/category.service';
import { QaAddComponent } from '../../components/qa-add/qa-add.component';
import { FilterQuestion, Question } from '../../interfaces/question';
import { QuestionService } from '../../repositories/question.service';
import { User } from '../../interfaces/user';
import { UserService } from '../../repositories/user.service';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';
import { MasterService } from 'src/app/services/master.service';

@Component({
  selector: 'app-qa',
  templateUrl: './qa.component.html',
  styleUrls: ['./qa.component.scss']
})

export class QaComponent implements OnInit {
  isFiltered = false;
  questionData: Question[] = [];
  categoryData: Category[] = [];
  userData: User[] = [];
  searchForm!: FormGroup;
  questionCount!: number;
  activeCollapse = false;

  constructor(
    private questionSvc: QuestionService,
    private categorySvc: CategoryService,
    private drawerSvc: NzDrawerService,
    private userSvc: UserService,
    private fb: FormBuilder,
    private translate: TranslateService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService,
    private masterSvc: MasterService
  ) { }

  ngOnInit(): void {
    const userId = localStorage.getItem('login_user_id') || ''
    this.getQuestionList(1, 10);
    this.searchForm = this.fb.group({
      language_ids: [[], [Validators.required]],
      tag_ids: [[], [Validators.required]],
      user_id: [[], [Validators.required]],
      title: [null, [Validators.required]],
    });
    this.masterSvc.userData$.subscribe((userData: User[]) => {
      this.userData = userData;
    });
    this.masterSvc.categoryData$.subscribe((categoryData: Category[]) => {
      this.categoryData = categoryData;
    });
  }

  getQuestionList(page: number, limit: number) {
    this.loaderSvc.call();
    this.questionSvc.getQuestionList(page, limit).subscribe({
      next: (questionListData: Question[]) => {
        this.questionData = questionListData;
        this.getQuestionCount();
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  getQuestionCount() {
    this.questionSvc.getQuestionCount().subscribe({
      next: (questionCount: { count: number }) => {
        this.questionCount = questionCount.count;
      },
      error: (error) => {
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  getFilteredQuestionCount(query: FilterQuestion) {
    this.questionSvc.getFilteredQuestionCount(query).subscribe({
      next: (filteredQuestionCount: { count: number }) => {
        this.questionCount = filteredQuestionCount.count;
      }
    })
  }

  openQaComponent(): void {
    const isMobile = window.innerWidth <= 768;
    const drawerWidth = isMobile ? '100vw' : '768px';
    const drawer = this.drawerSvc.create<QaAddComponent, { categoryData: Category[], userData: User[] }, string>({
      nzTitle: this.translate.instant('createQuestion'),
      nzPlacement: 'right',
      nzContent: QaAddComponent,
      nzMaskClosable: false,
      nzContentParams: {
        categoryData: this.categoryData,
        userData: this.userData
      },
      nzWidth: drawerWidth
    });
    drawer.afterClose.subscribe((data) => {
      if (data === 'created') {
        this.getQuestionList(1, 10);
      }
    });
  }

  filterQuestions(page: number, index: number): void {
    this.isFiltered = true;
    this.loaderSvc.call();
    this.questionSvc.searchQuestions(this.searchForm.value, page, index).subscribe({
      next: (filteredQuestions: Question[]) => {
        this.questionData = filteredQuestions;
        this.getFilteredQuestionCount(this.searchForm.value)
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.apiSvc.handleMessageError(error)
        this.apiSvc.handleErrorType(error)
      }
    })
  }

  clearFilterForms() {
    this.searchForm.controls['language_ids'].setValue([]);
    this.searchForm.controls['tag_ids'].setValue([]);
    this.searchForm.controls['user_id'].setValue([]);
    this.searchForm.controls['title'].setValue('');
  }

  categoryClicked(eventData: FilterQuestion) {
    this.searchForm.patchValue(eventData);
    this.getFilteredQuestionCount(eventData);
    this.activeCollapse = true;
  }
}
