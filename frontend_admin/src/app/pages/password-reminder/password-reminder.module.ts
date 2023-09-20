import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PasswordReminderRoutingModule } from './password-reminder-routing.module';
import { PasswordReminderComponent } from './password-reminder.component';
import { NgZorroAntdModule } from 'src/app/ng-zorro-antd/ng-zorro-antd.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';


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
