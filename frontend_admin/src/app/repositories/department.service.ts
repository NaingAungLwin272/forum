import { Injectable } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Observable } from 'rxjs';
import { Department } from '../interfaces/department';

@Injectable({
  providedIn: 'root'
})
export class DepartmentService extends ApiService {
  getDepartmentList(pageIndex?: number, pageSize?: number): Observable<Department[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/departments?page=${pageIndex}&limit=${pageSize}`,
      {}
    );
  }

  getDepartmentCount(): Observable<{ count: number }> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/departments/count`,
      {}
    )
  }

  createDepartment(formData: any): Observable<Department> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/department`,
      formData
    )
  }

  updateDepartment(id: string | undefined, formData: Department): Observable<Department> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/departments/${id}`,
      formData
    )
  }

  deleteDepartment(id: string | undefined): Observable<Department> {
    return this.apiConnecter(
      'DELETE',
      `${this.apiEndpoint}/departments/${id}`,
      {}
    )
  }
}
