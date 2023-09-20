import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UsersListRoutingModule } from './users-list-routing.module';
import { UsersListComponent } from './users-list.component';
import { SharedModule } from '../../shared.module';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { UserCreateComponent } from 'src/app/components/user-create/user-create.component';
import { ScrollingModule } from '@angular/cdk/scrolling';

@NgModule({
  declarations: [
    UsersListComponent,
    UserCreateComponent
  ],
  imports: [
    CommonModule,
    UsersListRoutingModule,
    SharedModule,
    ReactiveFormsModule,
    FormsModule,
    ScrollingModule
  ]
})
export class UsersListModule { }
