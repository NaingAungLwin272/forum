import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { QaDetailRoutingModule } from './qa-detail-routing.module';
import { QaDetailComponent } from './qa-detail.component';
import { NgZorroAntdModule } from '../../ng-zorro-antd/ng-zorro-antd.module';
import { SharedModule } from '../../shared.module';
import { TranslateModule } from '@ngx-translate/core';


@NgModule({
  declarations: [
    QaDetailComponent
  ],
  imports: [
    CommonModule,
    QaDetailRoutingModule,
    NgZorroAntdModule,
    SharedModule,
    TranslateModule
  ]
})
export class QaDetailModule { }
