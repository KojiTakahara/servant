var dir = angular.module('wxCardDirective', []);
dir.directive('cardlist', function() {
  return {
    restrict: 'E',
    templateUrl: '/view/common/cardlist.html',
    controller: function($scope) {
      $scope.setColor = function(color) {
        var result = '';
        switch (color){
        case 'white':
          result = '#F6BB42'; break;
        case 'red':
          result = '#DA4453'; break;
        case 'blue':
          result = '#3BAFDA'; break;
        case 'green':
          result = '#8CC152'; break;
        case 'black':
          result = '#967ADC'; break;
        case 'colorless':
          result = '#E6E9ED'; break;
        default:
          result = ''; break;
        }
        return "{backgroundColor: '" + result + "', borderColor: '" + result + "', height: '8px' }";
      };
    }
  };
});