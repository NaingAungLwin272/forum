import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '../../shared.module';
import { UserComponent } from './user.component';
import { UserRoutingModule } from './user-routing.module';
import { TranslateModule } from '@ngx-translate/core';
import { DatePipe } from '@angular/common';

@NgModule({
  declarations: [
    UserComponent
  ],
  imports: [
    CommonModule,
    UserRoutingModule,
    SharedModule,
    ReactiveFormsModule,
    TranslateModule
  ],
  providers: [DatePipe],
})
export class UserModule { }
