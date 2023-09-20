import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../services/api.service';
import { Department } from '../interfaces/department';

@Injectable({
  providedIn: 'root'
})
export class DepartmentService extends ApiService {


  getDepartment(id: string): Observable<Department> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/departments/${id}`,
      {}
    );
  }
}