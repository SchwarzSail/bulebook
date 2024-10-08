import 'package:blue_book/view/root-page/app_page.dart';
import 'package:blue_book/tansit_page.dart';
import 'package:blue_book/view/root-page/login_page.dart';

import 'app_routes.dart';
import 'package:get/get.dart';

class AppRouter {
  // 初始路由
  static const String initialRoute = AppRoutes.transit;
  // 404
  static final GetPage unknownRoute = GetPage(
    name: AppRoutes.notfound,
    page: () => AppPage(),
  );

  // 路由页面
  static final List<GetPage<dynamic>> getPages = [
    GetPage(
      name: AppRoutes.login,
      page:() => LoginPage(),
    ),
    GetPage(
      name: AppRoutes.transit,
      page: () => TansitPage(),
    ),
    GetPage(
      name: AppRoutes.app,
      page: () => AppPage(),
    ),
  ];
}
