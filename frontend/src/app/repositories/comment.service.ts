import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Comment, RequestComment } from '../interfaces/comment';
import { ApiService } from '../services/api.service';

@Injectable({
  providedIn: 'root'
})
export class CommentService extends ApiService {

  createComment(formData: RequestComment): Observable<Comment> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/comment`,
      formData
    )
  }

  updateComment(id: string, formData: RequestComment): Observable<Comment> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/comment/${id}`,
      formData
    )
  }

  getComment(id: string): Observable<Comment> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/comment/${id}`,
      {}
    )
  }

}
