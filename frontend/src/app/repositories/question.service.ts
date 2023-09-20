import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { FilterQuestion, Question, QuestionDetail } from '../interfaces/question';
import { ApiService } from '../services/api.service';

@Injectable({
  providedIn: 'root'
})
export class QuestionService extends ApiService {
  createQuestion(formData: Question): Observable<Question> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/question`,
      formData
    )
  }

  getQuestionList(page: number, limit: number, sort?: unknown, order?: string): Observable<Question[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/questions?page=${page}&limit=${limit}&sort=${sort}&order=${order}`,
      {}
    );
  }

  getQuestionById(id: string): Observable<QuestionDetail> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/questions/${id}`,
      {}
    );
  }

  getQuestionsWithUserId(userId: string, page: number, limit: number, sort?: unknown, order?: string): Observable<Question[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${userId}/questions?page=${page}&limit=${limit}&sort=${sort}&order=${order}`,
      {}
    )
  }

  searchQuestions(query: FilterQuestion, page?: number, limit?: number, sort?: unknown, order?: string): Observable<Question[]> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/questions/search?page=${page}&limit=${limit}&sort=${sort}&order=${order}`,
      { ...query }
    )
  }

  getQuestionCount(): Observable<{ count: number }> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/questions/count`,
      {}
    )
  }

  getFilteredQuestionCount(query: FilterQuestion, sort?: unknown, order?: string): Observable<{ count: number }> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/questions/filteredquestioncount?sort=${sort}&order=${order}`,
      { ...query }
    )
  }
}
