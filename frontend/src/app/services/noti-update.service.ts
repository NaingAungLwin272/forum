import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class NotiUpdateService {

  private notificationsMarkedAsReadSubject = new Subject<void>();

  notificationsMarkedAsRead$ = this.notificationsMarkedAsReadSubject.asObservable();

  notifyNotificationsMarkedAsRead() {
    this.notificationsMarkedAsReadSubject.next();
  }
}
