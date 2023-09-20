import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/qa' },
  { path: 'user/:id', canActivate: [AuthGuard], loadChildren: () => import('./pages/user/user.module').then(m => m.UserModule) },
  { path: 'qa', canActivate: [AuthGuard], loadChildren: () => import('./pages/qa/qa.module').then(m => m.QaModule) },
  { path: 'qa-detail/:id', canActivate: [AuthGuard], loadChildren: () => import('./pages/qa-detail/qa-detail.module').then(m => m.QaDetailModule) },
  { path: 'signin', loadChildren: () => import('./pages/signin/signin.module').then(m => m.SigninModule) },
  { path: 'password-reminder', loadChildren: () => import('./pages/password-reminder/password-reminder.module').then(m => m.PasswordReminderModule) },
  { path: 'password-reset', loadChildren: () => import('./pages/password-reset/password-reset.module').then(m => m.PasswordResetModule) },
  { path: '**', redirectTo: 'qa', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
