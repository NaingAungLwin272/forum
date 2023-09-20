import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { PasswordReminderRoutingModule } from './password-reminder-routing.module';
import { PasswordReminderComponent } from './password-reminder.component';
import { NgZorroAntdModule } from '../../ng-zorro-antd/ng-zorro-antd.module';


@NgModule({
  declarations: [
    PasswordReminderComponent
  ],
  imports: [
    CommonModule,
    PasswordReminderRoutingModule,
    NgZorroAntdModule,
    FormsModule, 
    ReactiveFormsModule
  ]
})
export class PasswordReminderModule { }
