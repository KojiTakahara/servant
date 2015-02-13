var fil = angular.module('trustFilter', ['ngSanitize']);
fil.filter('trustHtml', ['$sce', function($sce) {
  return function(html) {
    return $sce.trustAsHtml(html);
  };
}]);