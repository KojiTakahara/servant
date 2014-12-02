'use strict';

var app = angular.module('deckCtrl', [
  'apiService',
]);
app.controller('deckController', ['$scope', '$window', 'cardService', function($scope, $window, cardService) {
  cardService.getPublicDecks(100, 0).then(function(data) {
    $scope.decks = data;
  }, function() {
    // error
  });
}]);