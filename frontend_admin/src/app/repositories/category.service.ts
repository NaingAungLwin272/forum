import { Injectable } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Observable } from 'rxjs';
import { Category } from '../interfaces/category';


@Injectable({
  providedIn: 'root'
})
export class CategoryService extends ApiService {

  getCategoryList(pageIndex?: number, pageSize?: number): Observable<Category[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/categories?page=${pageIndex}&limit=${pageSize}`,
      {}
    );
  }

  getCategoryCount(): Observable<{ language_count: number, tag_count: number }> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/categories/count`,
      {}
    )
  }

  getCategoryByType(type?: number, pageIndex?: number, pageSize?: number): Observable<Category[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/categories/type?type=${type}&page=${pageIndex}&limit=${pageSize}`,
      {}
    );
  }

  updateCategory(id: string | undefined, formData: Category): Observable<Category> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/category/${id}`,
      formData
    )
  }

  deleteCategory(id: string | undefined): Observable<Category> {
    return this.apiConnecter(
      'DELETE',
      `${this.apiEndpoint}/category/${id}`,
      {}
    )
  }

  createCategory(formData: any): Observable<Category> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/category`,
      formData
    )
  }
}
