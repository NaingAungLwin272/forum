import { Injectable } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Observable } from 'rxjs';
import { FilterUser, User } from '../interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class UserService extends ApiService {
  getUsersList(page?: number, limit?: number): Observable<User[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users?page=${page}&limit=${limit}`,
      {}
    );
  }

  deleteUser(id: string): Observable<User> {
    return this.apiConnecter(
      'DELETE',
      `${this.apiEndpoint}/users/${id}`,
      {}
    )
  }

  updateUser(id: string, data: User): Observable<User> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/users/${id}`,
      { ...data }
    )
  }

  getUser(id: string): Observable<User> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/users/${id}`,
      {}
    )
  }

  createUser(data: User): Observable<User> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/user`,
      { ...data }
    )
  }

  createUserWithCsv(data: User[]): Observable<User[]> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/users/csv`,
      data
    )
  }

  filterUser(query: FilterUser, page?: number, limit?: number): Observable<User[]> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/users/search?page=${page}&limit=${limit}`,
      { ...query }
    )
  }

  getUserCount(data?: any) {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/users/count`,
      { ...data }
    )
  }
}
