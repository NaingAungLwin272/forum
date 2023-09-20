import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { PasswordReminderComponent } from './password-reminder.component';

const routes: Routes = [{ path: '', component: PasswordReminderComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PasswordReminderRoutingModule { }
