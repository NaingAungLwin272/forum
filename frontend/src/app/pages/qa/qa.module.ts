import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { QaRoutingModule } from './qa-routing.module';
import { QaComponent } from './qa.component';
import { SharedModule } from '../../shared.module';
import { QaAddComponent } from '../../components/qa-add/qa-add.component';
import { TranslateModule } from '@ngx-translate/core';
@NgModule({
  declarations: [
    QaComponent,
    QaAddComponent
  ],
  imports: [
    CommonModule,
    QaRoutingModule,
    SharedModule,
    TranslateModule
  ],
  exports: []
})
export class QaModule { }
