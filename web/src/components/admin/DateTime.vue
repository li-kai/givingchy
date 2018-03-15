<template>
  <time class="time">
    <span>{{formatted.weekday}}</span>
    <span>{{formatted.date}}</span>
    <span>{{formatted.time}}</span>
  </time>
</template>

<script>
const formatter = new Intl.DateTimeFormat('en-US', {
  weekday: 'long',
  year: 'numeric',
  month: 'short',
  day: 'numeric',
  hour: 'numeric',
  minute: 'numeric',
  timeZoneName: 'short',
});

export default {
  name: 'datetime',
  props: {
    datetime: {
      required: true,
    },
  },
  computed: {
    formatted() {
      const fullString = formatter.format(new Date(this.datetime));
      const [weekday, date, year, time] = fullString.split(', ');
      return {
        weekday,
        date: `${date}, ${year}`,
        time,
      };
    },
  },
};
</script>

<style scoped>
.time {
  display: flex;
  flex-direction: column;
}
</style>
