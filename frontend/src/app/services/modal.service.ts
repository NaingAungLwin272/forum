import { Injectable, EventEmitter } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ModalService {
  modalOpenedEmitter: EventEmitter<void> = new EventEmitter<void>();


  notifyModalOpened() {
    this.modalOpenedEmitter.emit();
  }
}
