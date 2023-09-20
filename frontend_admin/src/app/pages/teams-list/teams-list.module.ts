import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { TeamsListRoutingModule } from './teams-list-routing.module';
import { TeamsListComponent } from './teams-list.component';
import { SharedModule } from '../../shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    TeamsListComponent
  ],
  imports: [
    CommonModule,
    TeamsListRoutingModule,
    SharedModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class TeamsListModule { }
