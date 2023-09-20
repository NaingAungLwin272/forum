import { Injectable } from '@angular/core';
import { ApiService } from '../services/api.service';
import { Observable } from 'rxjs';
import { Noti } from '../interfaces/noti';


@Injectable({
  providedIn: 'root'
})
export class NotiService extends ApiService {
  createNoti(data: Noti): Observable<Noti> {
    return this.apiConnecter(
      'POST',
      `${this.apiEndpoint}/notification`,
      data
    )
  }

  getNoti(): Observable<Noti> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/notifications`,
      {}
    )
  }

  getNotiByUserId(userId: string, page: number, limit: number): Observable<Noti[]> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/user/${userId}/notifications?page=${page}&limit=${limit}`,
      {}
    )
  }

  getNotiCount(userId: string): Observable<{ count: number }> {
    return this.apiConnecter(
      'GET',
      `${this.apiEndpoint}/user/${userId}/noticount`,
      {}
    )
  }

  MarkAllNotiAsRead(userId: string): Observable<{ res: boolean }> {
    return this.apiConnecter(
      'PUT',
      `${this.apiEndpoint}/user/${userId}/notireset`,
      {}
    )
  }
}
