
export function convertCapitalCamelCaseToSnakeCase(str: string) {
  const newStr = str.replace(/([A-Z])/g, (match) => `_${match.toLowerCase()}`).slice(1)
  return newStr.replace(/,_/g, ',')
}
