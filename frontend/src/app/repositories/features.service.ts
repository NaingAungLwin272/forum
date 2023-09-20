import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../services/api.service';
import { Vote, CreateUserVote, UserVote, CreateView } from '../interfaces/vote';
import { User } from '../interfaces/user';
import { UserBadge, Badge } from '../interfaces/badge';
import { CreateBookmark, Bookmark } from '../interfaces/bookmark';

@Injectable({
  providedIn: 'root'
})
export class FeaturesService extends ApiService {

  getBookmarkByUserId(id: string, page: number, limit: number): Observable<Vote[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}/bookmarks?page=${page}&limit=${limit}`,
      {}
    );
  }

  getBookmarkByUserIdQuestionId(userId: string, questionId: string): Observable<Bookmark[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/bookmarks/user/${userId}/question/${questionId}`,
      {}
    )
  }

  createBookmark(formData: CreateBookmark): Observable<Bookmark> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/bookmarks`,
      formData
    )
  }

  deleteBookmark(id: string): Observable<Bookmark> {
    return this.apiConnecter(
      'DELETE',
      `${this.apiEndpoint}/bookmarks/${id}`,
      {}
    )
  }

  getVotesByUserId(id: string, page: number, limit: number): Observable<Vote[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}/votes?page=${page}&limit=${limit}`,
      {}
    )
  }

  getVotesByUserIdQuestionId(userId: string, questionId: string): Observable<UserVote[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/votes/user/${userId}/question/${questionId}`,
      {}
    )
  }

  createUserVote(formData: CreateUserVote): Observable<UserVote> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/votes`,
      formData
    )
  }

  deleteUserVote(id: string): Observable<UserVote> {
    return this.apiConnecter(
      'DELETE',
      `${this.apiEndpoint}/votes/${id}`,
      {}
    )
  }

  editProfile(id: string | undefined, formData: any): Observable<User> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/users/${id}`,
      formData
    )
  }

  getCommentByUserId(id: string, page: number, limit: number): Observable<Vote[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}/comments?page=${page}&limit=${limit}`,
      {}
    )
  }

  getAnswersByUserId(id: string, page: number, limit: number): Observable<Vote[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}/answers?page=${page}&limit=${limit}`,
      {}
    )
  }

  getBadgesByUserId(id: string): Observable<UserBadge[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}/badges`,
      {}
    )
  }

  getBadgeList(): Observable<Badge[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/badges`,
      {}
    )
  }

  getCommentByUserIdSolved(id: string, page: number, limit: number): Observable<Vote[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}/solved?page=${page}&limit=${limit}`,
      {}
    )
  }

  getMentionByUserId(id: string, page: number, limit: number): Observable<Vote[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}/mentions?page=${page}&limit=${limit}`,
      {}
    )
  }

  getBookMarkByUserIdQuestionId(user_id: string, question_id: string): Observable<Bookmark[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/bookmarks/user/${user_id}/question/${question_id}`,
      {}
    )
  }
  createMention(formData: CreateBookmark): Observable<Bookmark> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/mentions`,
      formData
    )
  }

  createViewByUserIdQuestionId(data: CreateView): Observable<CreateView> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/views`,
      { ...data }
    )
  }
}
