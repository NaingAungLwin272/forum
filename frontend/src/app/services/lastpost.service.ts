import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class LastpostService {

  private lastpostSubject = new BehaviorSubject<string>('');

  lastpost$ = this.lastpostSubject.asObservable();

  setLastpost(lastpost: string) {
    this.lastpostSubject.next(lastpost);
  }

  getLastpost() {
    return this.lastpostSubject.getValue();
  }
}
