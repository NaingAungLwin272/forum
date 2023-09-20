import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../services/api.service';
import { Category } from '../interfaces/category';

@Injectable({
  providedIn: 'root'
})
export class CategoryService extends ApiService {

  getCategoryList(): Observable<Category[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/categories`,
      {}
    );
  }

  getCategory(id: string[]): Observable<Category> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/category/${id}`,
      {}
    )
  }
}
