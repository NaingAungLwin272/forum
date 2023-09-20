import { RouterModule, Routes } from '@angular/router';

import { AuthGuard } from './guards/auth.guard';
import { NgModule } from '@angular/core';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/user-list' },
  { path: 'user-list', canActivate: [AuthGuard], loadChildren: () => import('./pages/users-list/users-list.module').then(m => m.UsersListModule) },
  { path: 'signin', loadChildren: () => import('./pages/signin/signin.module').then(m => m.SigninModule) },
  { path: 'password-reset', loadChildren: () => import('./pages/password-reset/password-reset.module').then(m => m.PasswordResetModule) },
  { path: 'password-reminder', loadChildren: () => import('./pages/password-reminder/password-reminder.module').then(m => m.PasswordReminderModule) },
  { path: 'department-list', canActivate: [AuthGuard], loadChildren: () => import('./pages/departments-list/departments-list.module').then(m => m.DepartmentsListModule) },
  { path: 'team-list', canActivate: [AuthGuard], loadChildren: () => import('./pages/teams-list/teams-list.module').then(m => m.TeamsListModule) },
  { path: 'category-list', canActivate: [AuthGuard], loadChildren: () => import('./pages/categories-list/categories-list.module').then(m => m.CategoriesListModule) },
  { path: 'message', canActivate: [AuthGuard], loadChildren: () => import('./pages/messages-list/messages-list.module').then(m => m.MessagesListModule) },
  { path: '**', redirectTo: 'user-list', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
