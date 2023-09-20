import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { QaDetailComponent } from './qa-detail.component';

const routes: Routes = [{ path: '', component: QaDetailComponent }];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class QaDetailRoutingModule { }
