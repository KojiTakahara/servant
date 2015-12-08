var service = angular.module('amazonService', []);
service.factory('amazonService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.search = function() {
    var deferred = $q.defer();
    $http({
      method: 'GET',
      url: '/api/amazon',
      cache: true
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  service.create = function(data) {
    var deferred = $q.defer();
    $http({
      method: 'POST',
      url: '/api/amazon',
      data: data
    }).success(function(data, status, headers, config) {
      deferred.resolve(data);
    }).error(function(data, status, headers, config) {
      deferred.reject(data);
    });
    return deferred.promise;
  };

  return service;
}]);