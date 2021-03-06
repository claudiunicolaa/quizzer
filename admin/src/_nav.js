export default {
  items: [
    {
      name: 'Dashboard',
      url: '/dashboard',
      icon: 'icon-speedometer',
      badge: {
        variant: 'info',
        text: 'NEW',
      },
    },
      {
          title: true,
          name: 'Quiz',
          wrapper: {
              element: '',
              attributes: {},
          },
      },
      {
          name: 'Question templates',
          url: '/quiz/question-templates',
          icon: 'icon-puzzle',
      },
      {
          name: 'Quiz templates',
          url: '/quiz/quiz-templates',
          icon: 'icon-calculator',
      },
      {
          title: true,
          name: 'User',
          wrapper: {
              element: '',
              attributes: {},
          },
      },
      {
          name: 'Management',
          url: '/users',
          icon: 'icon-people',
      },
      {
          name: 'Results',
          url: '/quiz/results',
          icon: 'icon-speedometer',
      },
  ],
};
