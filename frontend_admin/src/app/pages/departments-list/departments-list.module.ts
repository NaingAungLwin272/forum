import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DepartmentsListRoutingModule } from './departments-list-routing.module';
import { DepartmentsListComponent } from './departments-list.component';
import { SharedModule } from '../../shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    DepartmentsListComponent
  ],
  imports: [
    CommonModule,
    DepartmentsListRoutingModule,
    SharedModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class DepartmentsListModule { }
