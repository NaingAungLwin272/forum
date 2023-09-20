import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { NgZorroAntdModule } from 'src/app/ng-zorro-antd/ng-zorro-antd.module';
import { PasswordResetComponent } from './password-reset.component';
import { PasswordResetRoutingModule } from './password-reset-routing.module';

@NgModule({
  declarations: [
    PasswordResetComponent
  ],
  imports: [
    CommonModule,
    PasswordResetRoutingModule,
    NgZorroAntdModule,
    FormsModule, 
    ReactiveFormsModule
  ]
})
export class PasswordResetModule { }
