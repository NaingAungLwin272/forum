import { Injectable } from '@angular/core';
import { UserService } from '../repositories/user.service';
import { CategoryService } from '../repositories/category.service';
import { User } from '../interfaces/user';
import { Category } from '../interfaces/category';
import { BehaviorSubject } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class MasterService {

  public _userData = new BehaviorSubject<User[]>([]);
  public userData$ = this._userData.asObservable();

  public _categoryData = new BehaviorSubject<Category[]>([]);
  public categoryData$ = this._categoryData.asObservable();

  userData: User[] = [];
  categoryData: Category[] = [];

  constructor(
    private userSvc: UserService,
    private categorySvc: CategoryService,
  ) {}

  getUserList() {
    return new Promise(resolve => {
      this.userSvc.getUserList().subscribe({
        next: (userData: User[]) => {
          this._userData.next(userData);
          resolve(userData);
        },
        error: (error) => {
          resolve(error);
        }
      });
    });
  }

  getCategoryList() {
    return new Promise(resolve => {
      this.categorySvc.getCategoryList().subscribe({
        next: (categoryData: Category[]) => {
          this._categoryData.next(categoryData);
          resolve(categoryData);
        },
        error: (error) => {
          resolve(error);
        }
      });
    });

  }
}
