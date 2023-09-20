import { Injectable } from '@angular/core';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal';
import { LoadingComponent } from '../components/loading/loading.component';

@Injectable({
  providedIn: 'root'
})
export class LoaderService {
  private loaderModalRef: NzModalRef | null = null;
  constructor(public modal: NzModalService) { }

  call() {
    this.loaderModalRef = this.modal.create({
      nzContent: LoadingComponent,
      nzFooter: null,
      nzClosable: false,
      nzCentered: true,
      nzStyle: { 'width': '150px' },
      nzNoAnimation: true,
      nzMaskClosable: false
    });
  }

  dismiss() {
    if (this.loaderModalRef) {
      this.loaderModalRef.close();
      this.loaderModalRef = null;
    }
  }

  dissAll() {
    this.modal.closeAll()
  }
}
