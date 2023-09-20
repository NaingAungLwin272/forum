import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Category } from 'src/app/interfaces/category';
import { FilterQuestion, Question } from 'src/app/interfaces/question';
import { QuestionService } from 'src/app/repositories/question.service';
import { FeaturesService } from 'src/app/repositories/features.service';
import { CreateView } from 'src/app/interfaces/vote';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';
import { Router } from '@angular/router';
import { User } from 'src/app/interfaces/user';
import { MasterService } from 'src/app/services/master.service';

@Component({
  selector: 'app-question-table',
  templateUrl: './question-table.component.html',
  styleUrls: ['./question-table.component.scss']
})

export class QuestionTableComponent {
  currentPage = 1;
  //temp
  userId = localStorage.getItem('login_user_id');
  sortClicked = false;
  sort !: unknown;
  order !: string;
  userData: User[] = [];
  @Output() categoryClicked: EventEmitter<FilterQuestion> = new EventEmitter<FilterQuestion>();
  @Input() data: Question[] = [];
  @Input() type!: string;
  @Input() categoryData: Category[] = [];
  @Input() isFiltered!: boolean;
  @Input() filteredFormData!: FilterQuestion;
  @Output() indexChange = new EventEmitter<number>()
  @Input() summaryCount!: number;
  @Input() filteredQuestionCount !: number;
  isLoading = false;
  darkTheme!: boolean;
  uniqueUserIds = new Set<string>();
  commentUserIds: Map<string, Set<string>> = new Map();
  commentList!: Comment[];

  constructor(
    private questionSvc: QuestionService,
    private featureSvc: FeaturesService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService,
    private router: Router,
    private masterSvc: MasterService,
  ) {
    this.masterSvc.userData$.subscribe((userData: User[]) => {
      this.userData = userData;
    });
  }

  loadMore() {
    this.isLoading = true;
    const limit = 10;
    if (!this.isFiltered || !this.sortClicked) {
      this.currentPage = 1;
      this.currentPage++;
    }
    if (this.sortClicked) {
      this.currentPage = 1;
      this.currentPage++;
    }
    if (this.isFiltered) {
      this.currentPage = 1;
      this.currentPage++;
    }
    if (this.type === 'user') {
      if (this.userId != null) {
        this.questionSvc.getQuestionsWithUserId(this.userId, this.currentPage, limit, this.sort, this.order).subscribe({
          next: (data: Question[]) => {
            this.data = [...this.data, ...data];
            this.isLoading = false;
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
            this.isLoading = false;
          }
        });
      }
    } else if (this.isFiltered) {
      this.questionSvc.searchQuestions(this.filteredFormData, this.currentPage, limit, this.sort, this.order).subscribe({
        next: (data: Question[]) => {
          this.data = [...this.data, ...data];
          this.questionSvc.getFilteredQuestionCount(this.filteredFormData, this.sort, this.order).subscribe({
            next: (data) => {
              this.summaryCount = data.count;
              this.isLoading = false;
            }
          })
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
          this.isLoading = false;
        }
      });
    } else if (this.sortClicked) {
      this.questionSvc.getQuestionList(this.currentPage, limit, this.sort, this.order).subscribe({
        next: (sortedData: Question[]) => {
          this.data = [...this.data, ...sortedData];
          this.isLoading = false;
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
          this.isLoading = false;
        }
      })
    }
    else {
      this.questionSvc.getQuestionList(this.currentPage, limit).subscribe({
        next: (data: Question[]) => {
          this.data = [...this.data, ...data];
          this.isLoading = false;
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
          this.isLoading = false;
        }
      });
    }
  }

  getCategoryName(id: string): string {
    return this.categoryData.find((category: Category) => category._id === id)?.name || '';
  }

  onSortChange(order: string, sort: unknown): void {
    this.sort = sort;
    this.order = order;
    this.sortClicked = true;
    if (this.type === 'user') {
      this.loaderSvc.call();
      if (this.userId != null) {
        this.questionSvc.getQuestionsWithUserId(this.userId, 1, 10, sort, order).subscribe({
          next: (data: Question[]) => {
            this.data = data;
            this.loaderSvc.dismiss();
          },
          error: (error) => {
            this.apiSvc.handleErrorType(error);
          }
        });
      }
    }
    else if (this.isFiltered) {
      this.loaderSvc.call();
      this.questionSvc.searchQuestions(this.filteredFormData, 1, 10, sort, order).subscribe({
        next: (data: Question[]) => {
          this.data = data;
          this.loaderSvc.dismiss();
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
        }
      });
    }
    else {
      this.loaderSvc.call();
      this.questionSvc.getQuestionList(1, 10, sort, order).subscribe({
        next: (sortedData: Question[]) => {
          this.data = sortedData;
          this.loaderSvc.dismiss();
        },
        error: (error) => {
          this.apiSvc.handleErrorType(error);
        }
      })
    }
  }

  createViewByUserIdQuestionId(user_id: string | null, question_id: string): void {
    if (user_id != null) {
      const data: CreateView = {
        user_id: user_id,
        question_id: question_id,
      }
      this.loaderSvc.call();
      this.featureSvc.createViewByUserIdQuestionId(data).subscribe({
        next: () => {
          this.loaderSvc.dismiss();
        },
        error: () => {
          this.loaderSvc.dismiss();
        }
      });
    }
  }

  filterByTag(tag: string, desc: string): void {
    if (this.router.url !== '/qa') {
      return;
    }
    this.loaderSvc.call();
    this.filteredFormData.language_ids = [];
    this.filteredFormData.tag_ids = [];
    if (desc === 'language') {
      this.filteredFormData.language_ids.push(tag);
      this.categoryClicked.emit(this.filteredFormData)
    }
    else {
      this.filteredFormData.tag_ids.push(tag);
      this.categoryClicked.emit(this.filteredFormData)
    }
    this.questionSvc.searchQuestions(this.filteredFormData, 1, 10).subscribe({
      next: (data: Question[]) => {
        this.data = data;
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.apiSvc.handleErrorType(error);
      }
    });
  }

  getUserImageLink(id: string): string {
    return this.userData.find((user: User) => user._id === id)?.profile || '../../../assets/images/avatar.png';
  }

  getUserName(id: string): string {
    return this.userData.find((user: User) => user._id === id)?.display_name || '';
  }
}
