export function formatDate(date: Date): string {
    const str = date.toISOString();
    return str.substr(0, 10);
}

export function formatTomorrowDate(date: Date): string {
    date.setDate(date.getDate() + 1);
    return formatDate(date);
}
