import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MessagesListRoutingModule } from './messages-list-routing.module';
import { MessagesListComponent } from './messages-list.component';
import { SharedModule } from '../../shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';



@NgModule({
  declarations: [
    MessagesListComponent
  ],
  imports: [
    CommonModule,
    MessagesListRoutingModule,
    SharedModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class MessagesListModule { }
