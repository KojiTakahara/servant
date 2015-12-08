var service = angular.module('userService', []);
service.factory('userService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.getLoginUser = function() {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/loginUser'
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  return service;
}]);